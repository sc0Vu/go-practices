package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var (
	// API - Api general config
	API APIConf

	// DB - DB general config
	DB DBConf
)

// APIConf config
type APIConf struct {
	APIprotocol string   `yaml:"protocol"`
	Domain      string   `yaml:"domain"`
	Port        string   `yaml:"port"`
	Debug       bool     `yaml:"debug"`
	CORSDomains []string `yaml:"cors_domains"`
}

// DBConf config
type DBConf struct {
	Table       string `yaml:"table"`
	Host        string `yaml:"host"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Port        int    `yaml:"port"`
	MaxConn     int    `yaml:"max_connections"`
	MaxIdleConn int    `yaml:"max_idle_connections"`
}

// DeploySet - deploy config
type DeploySet struct {
	API APIConf `yaml:"api"`
	DB  DBConf  `yaml:"db"`
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
	DB = dconf.DB
}
