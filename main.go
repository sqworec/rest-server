package main

import (
	"rest-server/api"
	"rest-server/data"
	"rest-server/service"

	"github.com/jinzhu/configor"
)

var Config AppConfig

func main() {
	configor.New(nil).Load(&Config, "config.yaml")

	dictionary := data.NewDictionary(Config.DB)

	service := service.NewService(dictionary)

	api := api.NewAPI(Config.API, service)

	api.StartServer()
}