package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	_ "github.com/lib/pq"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
	Server struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"server"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

var db *sql.DB

func main() {

	configFile, openErr := os.Open("config.json")
	if openErr != nil {
		log.Fatal("Error opening config file:", openErr)
	}
	defer configFile.Close()

	var config Config
	decodeErr := json.NewDecoder(configFile).Decode(&config)
	if decodeErr != nil {
		log.Fatal("Error decoding config JSON:", decodeErr)
	}

	dbInfo := "host=" + config.Database.Host +
		" port=" + strconv.Itoa(config.Database.Port) +
		" user=" + config.Database.User +
		" password=" + config.Database.Password +
		" dbname=" + config.Database.DBName +
		" sslmode=" + config.Database.SSLMode

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	consulConfig := api.DefaultConfig()
	consulConfig.Address = "consul:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("Failed to create Consul client:", err)
		return
	}

	err = registerService(consulClient, config.Server.Address, config.Server.Port, "/health")
	if err != nil {
		fmt.Println("Failed to register service :", err)
		return
	}

	defer func() {
		err := deregisterService(consulClient, config.Server.Address)
		if err != nil {
			fmt.Println("Failed to deregister service :", err)
		}
	}()

	router := gin.Default()
	router.GET("/user/users", getUsers)
	router.GET("/user/:id", getUserById)
	router.POST("/user/create", createUser)
	router.PUT("/user/update/:id", updateUser)
	router.DELETE("/user/delete/:id", deleteUser)

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// serverAddress := config.Server.Address + ":" + strconv.Itoa(config.Server.Port)
	serverAddress := ":" + strconv.Itoa(config.Server.Port)
	router.Run(serverAddress)

	ticker := time.NewTicker(300 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			ok, err := checkHealth(consulClient, config.Server.Address)
			if err != nil {
				fmt.Println("Failed to check health:", err)
				continue
			}
			if !ok {
				err := deregisterService(consulClient, config.Server.Address)
				if err != nil {
					fmt.Println("Failed to deregister service :", err)
					continue
				}
				err = registerService(consulClient, config.Server.Address, config.Server.Port, "/health")
				if err != nil {
					fmt.Println("Failed to register service :", err)
				}
			}
		}
	}
}

func registerService(client *api.Client, serviceName string, servicePort int, healthCheckPath string) error {
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
		return fmt.Errorf("Failed to register %s service: %v", serviceName, err)
	}

	fmt.Printf("Registered %s service\n", serviceName)
	return nil
}

func deregisterService(client *api.Client, serviceName string) error {
	err := client.Agent().ServiceDeregister(serviceName)
	if err != nil {
		return fmt.Errorf("Failed to deregister %s service: %v", serviceName, err)
	}

	fmt.Printf("Deregistered %s service\n", serviceName)
	return nil
}

func checkHealth(client *api.Client, serviceName string) (bool, error) {
	healthChecks, _, err := client.Health().Checks(serviceName, nil)
	if err != nil {
		return false, fmt.Errorf("Failed to query health checks: %v", err)
	}

	for _, check := range healthChecks {
		if check.Status != api.HealthPassing {
			return false, nil
		}
	}

	return true, nil
}

func getUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, username, fullname FROM userdata")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var a User
		err := rows.Scan(&a.ID, &a.Username, &a.Fullname)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")

	rows, err := db.Query("SELECT id, username, fullname FROM userdata WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Fullname)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {

	var userAlbum User
	if err := c.BindJSON(&userAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO userdata (id, username, fullname) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userAlbum.ID, userAlbum.Username, userAlbum.Fullname); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, userAlbum)
}

func updateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare("UPDATE userdata SET username=$2, fullname=$3 WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id, user.Username, user.Fullname)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
		return
	}

	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func deleteUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")

	stmt, err := db.Prepare("DELETE FROM userdata WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
		return
	}

	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
