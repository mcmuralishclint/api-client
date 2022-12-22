package client

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	services []Service
}

type Service struct {
	base     string
	versions []Version
}

type Version struct {
	apis []Api
}
type Api struct {
	endpoint string
	verb     string
}

var config *Config

func Init() {
	c := &Config{}
	config = c.loadConfig()
}

func (c *Config) loadConfig() *Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return c
}

func Test() {
	fmt.Println(&config)
	fmt.Println(config)
}
