package main

var configuration *Config

func init() {
	configuration = parser("config.yml")
}

func main() {
	req := &Req{Service: "product", Version: "v1", Path: "get_products"}
	req.MakeHttpCall()
}
