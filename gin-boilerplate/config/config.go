package config

import (
	"io/ioutil"
	"log"

	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

var (
	// API - Api general config
	API APIConf
)

// APIConf config
type APIConf struct {
	APIprotocol string   `yaml:"api_protocol"`
	Domain      string   `yaml:"api_domain"`
	Port        string   `yaml:"api_port"`
	Debug       bool     `yaml:"api_debug"`
	CORSDomains []string `yaml:"api_cors_domains"`
}

// DeploySet - deploy config
type DeploySet struct {
	API      APIConf  `yaml:"api"`
}

// ReadConf read the config file from input filePath
func init() {
	var dconf DeploySet
	if content, ioErr := ioutil.ReadFile("./config/conf.yaml"); ioErr != nil {
		logrus.Fatalf("read service config file error: %v", ioErr)
	} else {
		if ymlErr := yaml.Unmarshal(content, &dconf); ymlErr != nil {
			log.Fatalf("error while unmarshal from db config: %v", ymlErr)
		}
	}

	API = dconf.API
}
