package consul

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"backend/user-service/config"

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

func RegisterService(client *api.Client, serviceName string, servicePort int, healthCheckPath string) error {
	registration := &api.AgentServiceRegistration{
		Name:    serviceName,
		Address: serviceName,
		Port:    servicePort,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d%s", serviceName, servicePort, healthCheckPath),
			Interval:                       "1m",
			Timeout:                        "10s",
			DeregisterCriticalServiceAfter: "3m",
		},
	}

	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		return fmt.Errorf("failed to register %s service: %v", serviceName, err)
	}

	fmt.Printf("Registered %s service\n", serviceName)
	return nil
}

func DeregisterService(client *api.Client, serviceName string) error {
	err := client.Agent().ServiceDeregister(serviceName)
	if err != nil {
		return fmt.Errorf("failed to deregister %s service: %v", serviceName, err)
	}

	fmt.Printf("Deregistered %s service\n", serviceName)
	return nil
}

func CheckHealth(client *api.Client, serviceName string) (bool, error) {
	healthChecks, _, err := client.Health().Checks(serviceName, nil)
	if err != nil {
		return false, fmt.Errorf("failed to query health checks: %v", err)
	}

	for _, check := range healthChecks {
		if check.Status != api.HealthPassing {
			return false, nil
		}
	}

	return true, nil
}

func RegisterWithHealthCheck(consulClient *api.Client, serviceName string, servicePort int, healthCheckPath string, checkInterval time.Duration) {
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for range ticker.C {
		ok, err := CheckHealth(consulClient, serviceName)
		if err != nil {
			log.Fatal("Failed to check health:", err)
			continue
		}
		if !ok {
			err := DeregisterService(consulClient, serviceName)
			if err != nil {
				log.Fatal("Failed to deregister service:", err)
				continue
			}
			err = RegisterService(consulClient, serviceName, servicePort, healthCheckPath)
			if err != nil {
				log.Fatal("Failed to register service:", err)
			}
		}
	}
}
