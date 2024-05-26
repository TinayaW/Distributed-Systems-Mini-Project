package router

import (
	"net/http"
	"net/http/httputil"
	"strconv"

	"backend/api-gateway/internal/consul"
	"backend/api-gateway/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func SetupRouter(consulClient *api.Client) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	router.NoRoute(func(c *gin.Context) {
		serviceName, err := utils.GetServiceName(c.Request.URL.Path)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL path"})
			return
		}

		service, err := consul.GetService(consulClient, serviceName)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service not available"})
			return
		}

		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = service.Address + ":" + strconv.Itoa(service.Port)
			req.URL.Path = c.Request.URL.Path
		}}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	return router
}
