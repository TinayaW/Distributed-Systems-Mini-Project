package main

import (
	"log"
	"strconv"

	"backend/api-gateway/config"
	"backend/api-gateway/internal/consul"
	"backend/api-gateway/internal/router"

	"github.com/hashicorp/consul/api"
)

func main() {
	configdata, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
		return
	}

	consulClient, err := consul.CreateConsulClient(configdata)
	if err != nil {
		log.Fatal("Failed to create Consul client:", err)
		return
	}

	startServer(configdata, consulClient)
}

func startServer(configdata *config.Config, consulClient *api.Client) {
	router := router.SetupRouter(consulClient)

	serverAddress := configdata.Endpoint.Address + ":" + strconv.Itoa(configdata.Endpoint.Port)
	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("Failed to start server:", err)
		return
	}
}
