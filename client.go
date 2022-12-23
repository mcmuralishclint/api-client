package main

import "fmt"

var configuration *Config

type Req struct {
	Service string
	Version string
	Name    string
}

type Res struct {
	Path string
	Verb string
}

func init() {
	configuration = parser("config.yml")
}

func main() {
	req := &Req{Service: "laas", Version: "v1", Name: "create_shipment"}
	res := req.RetreiveConfig()
	fmt.Println(res)
}

func (request *Req) RetreiveConfig() *Res {
	service := configuration.Services[request.Service]
	baseUrl := service.BaseUrl
	api := service.Version[request.Version][request.Name]
	endpoint := baseUrl + api.Path
	return &Res{Path: endpoint, Verb: api.Verb}
}
