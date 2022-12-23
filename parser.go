package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Info     interface{} `yaml:"info"`
	Services map[string]Service
}

type Service struct {
	Name    string `yaml:"name"`
	BaseUrl string `yaml:"base_url"`
	Version map[string]map[string]Endpoint
}

type Endpoint struct {
	Path string
	Verb string
}

func parser(filename string) *Config {
	ymlData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	configs := Config{}
	yaml.Unmarshal(ymlData, &configs)
	return &configs
}
