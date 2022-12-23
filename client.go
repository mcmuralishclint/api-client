package main

import "fmt"

var configuration *Config

func init() {
	configuration = parser("config.yml")
}

func main() {
	req := &Req{Service: "laas", Version: "v1", Name: "create_shipment"}
	res := req.RetreiveConfig()
	fmt.Println(res)
}
