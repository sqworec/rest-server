package main

import (
	"rest-server/api"
	"rest-server/data"
)

type AppConfig struct {
	API api.APIConfig
	DB  data.DBConfig
}
