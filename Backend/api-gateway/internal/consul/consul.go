package consul

import (
	"fmt"
	"strconv"

	"backend/api-gateway/config"

	"github.com/hashicorp/consul/api"
)

func CreateConsulClient(configdata *config.Config) (*api.Client, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = configdata.Consul.Address + ":" + strconv.Itoa(configdata.Consul.Port)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Consul client: %v", err)
	}
	return consulClient, nil
}

func GetService(client *api.Client, serviceName string) (*api.AgentService, error) {
	passingOnly := true
	services, _, err := client.Health().Service(serviceName, "", passingOnly, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query service: %v", err)
	}
	if len(services) == 0 {
		return nil, fmt.Errorf("service not found: %s", serviceName)
	}

	agentService := &api.AgentService{
		ID:      services[0].Service.ID,
		Service: services[0].Service.Service,
		Address: services[0].Service.Address,
		Port:    services[0].Service.Port,
	}
	return agentService, nil
}
