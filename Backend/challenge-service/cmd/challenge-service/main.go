package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"backend/challenge-service/config"
	"backend/challenge-service/internal/consul"
	"backend/challenge-service/internal/postgre"
	"backend/challenge-service/internal/router"

	"github.com/hashicorp/consul/api"
)

func main() {
	configdata, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
		return
	}

	db, err := postgre.NewDBConnection(configdata)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
		return
	}

	consulClient, err := consul.CreateConsulClient(configdata)
	if err != nil {
		log.Fatal("Failed to create Consul client:", err)
		return
	}

	err = consul.RegisterService(consulClient, configdata.Endpoint.Name, configdata.Endpoint.Port, "/health")
	if err != nil {
		log.Fatal("Failed to register service :", err)
		return
	}

	defer func() {
		err := consul.DeregisterService(consulClient, configdata.Endpoint.Name)
		if err != nil {
			log.Fatal("Failed to deregister service :", err)
		}
	}()

	startServer(configdata, consulClient, db)

	consul.RegisterWithHealthCheck(consulClient, configdata.Endpoint.Name, configdata.Endpoint.Port, "/health", 300*time.Second)
}

func startServer(configdata *config.Config, consulClient *api.Client, db *sql.DB) {
	router := router.SetupRouter(consulClient, db)

	serverAddress := configdata.Endpoint.Address + ":" + strconv.Itoa(configdata.Endpoint.Port)
	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("Failed to start server:", err)
		return
	}
}
