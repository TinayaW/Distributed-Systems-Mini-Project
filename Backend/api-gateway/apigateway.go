package main

import (
	"encoding/json"
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Service struct {
		UserService       string `json:"user_service"`
		ChallengeService  string `json:"challenge_service"`
		SubmissionService string `json:"submission_service"`
	} `json:"service"`
	Server struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"server"`
}

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

	router := gin.Default()

	router.Any("/user/*action", func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   config.Service.UserService,
		})
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	router.Any("/challenge/*action", func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   config.Service.ChallengeService,
		})
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	router.Any("/submission/*action", func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   config.Service.SubmissionService,
		})
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	serverAddress := config.Server.Address + ":" + strconv.Itoa(config.Server.Port)
	router.Run(serverAddress)
}
