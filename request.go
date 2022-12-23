package main

type Req struct {
	Service string
	Version string
	Name    string
}

type Res struct {
	Path string
	Verb string
}

func (request *Req) RetreiveConfig() *Res {
	service := configuration.Services[request.Service]
	baseUrl := service.BaseUrl
	api := service.Version[request.Version][request.Name]
	endpoint := baseUrl + api.Path
	return &Res{Path: endpoint, Verb: api.Verb}
}
