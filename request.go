package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var muxRouter = mux.NewRouter()

type Req struct {
	Service string
	Version string
	Path    string
}

type Res struct {
	Path string
	Verb string
}

func (request *Req) RetreiveConfig() *Res {
	service := configuration.Services[request.Service]
	baseUrl := service.BaseUrl
	api := service.Version[request.Version][request.Path]
	endpoint := baseUrl + api.Path
	return &Res{Path: endpoint, Verb: api.Verb}
}

func (request *Req) MakeHttpCall() string {
	res := request.RetreiveConfig()
	fmt.Println(res)
	if res.Verb == "GET" {
		response, err := http.Get(res.Path)
		if err != nil {
			fmt.Printf("Error when querying the data for %s", request.Path, err)
		}
		body, err := ioutil.ReadAll(response.Body)
		return string(body)
	}

	return ""
}
