package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/op/go-logging"
	"github.com/zenazn/goji"
)

type ApiConfiguration struct {
	QueryApiRootName string
}

type DnsConfiguration struct {
	TrustedDnsServers []string
}

type Config struct {
	ApiConfig ApiConfiguration
	DnsConfig DnsConfiguration
}

var log *logging.Logger = logging.MustGetLogger("doh")
var config Config = readConfiguration()

func readConfiguration() Config {
	tomlData, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Error(err)
		panic(err)
	}

	var config Config
	_, err = toml.Decode(string(tomlData), &config)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	return config
}

func configureRoutes(config ApiConfiguration) {
	apiRoot := "/" + config.QueryApiRootName
	goji.Get(apiRoot+"/:hostname", query)
}

func main() {
	log.Debugf("Read configuration: %+v", config)

	log.Info("Configuring REST API")
	configureRoutes(config.ApiConfig)

	log.Info("Starting REST API")
	goji.Serve()
}
