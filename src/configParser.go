package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// TMConfig holds the configuration for the entire TimberMill application
type TMConfig struct {
	Frontend struct {
		Enabled bool   `yaml:"enabled"`
		Listen  string `yaml:"listen"`
		Ssl     struct {
			Enabled     bool   `yaml:"enabled"`
			Certificate string `yaml:"certificate"`
			Privatekey  string `yaml:"privatekey"`
		} `yaml:"ssl"`
	} `yaml:"frontend"`

	Collectors []struct {
		Name     string `yaml:"name"`
		Type     string `yaml:"type"`
		Protocol string `yaml:"protocol"`
		Port     uint16 `yaml:"port"`
		Storage  string `yaml:"storage"`
	} `yaml:"collectors"`

	Storage []struct {
		Name string
		Type string
		URL  string
	} `yaml:"storage"`
}

// LoadConfig loads the TimberMill configuration from a YAML file
func LoadConfig(configPath string) *TMConfig {
	confstr, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicln("Unable to read configuration file", configPath)
	}
	var conf TMConfig
	err = yaml.Unmarshal(confstr, &conf)
	if err != nil {
		log.Panic("Unable to parse configuration file", configPath)
	}
	return &conf
}
