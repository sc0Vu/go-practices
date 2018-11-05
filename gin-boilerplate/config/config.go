package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var (
	// API - Api general config
	API APIConf
)

// APIConf config
type APIConf struct {
	APIprotocol string   `yaml:"protocol"`
	Domain      string   `yaml:"domain"`
	Port        string   `yaml:"port"`
	Debug       bool     `yaml:"debug"`
	CORSDomains []string `yaml:"cors_domains"`
}

// DeploySet - deploy config
type DeploySet struct {
	API APIConf `yaml:"api"`
}

// ReadConf read the config file from input filePath
func init() {
	var dconf DeploySet
	if content, ioErr := ioutil.ReadFile("./config/conf.yaml"); ioErr != nil {
		log.Fatalf("read service config file error: %v", ioErr)
	} else {
		if ymlErr := yaml.Unmarshal(content, &dconf); ymlErr != nil {
			log.Fatalf("error while unmarshal from db config: %v", ymlErr)
		}
	}

	API = dconf.API
}
