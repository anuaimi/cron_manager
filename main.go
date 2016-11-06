package main

// cron manager

import (
	"flag"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {

	// read parameters
	port := os.Getenv("CM_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	
	// get input parameters
	apiKey := flag.String("api_key", os.Getenv("CM_APIKEY"), "api key to authenticate")
	apiPort := flag.String("port", port,"port for API")
	flag.Parse()

	// setup api routes
	r := gin.Default()
	authorized := r.Group("/")
	if len(*apiKey) > 0 {
		authorized = r.Group("/", gin.BasicAuth(gin.Accounts{"apiKey" : *apiKey}))
	}

	authorized.POST("/job", submit_job)

	r.Run(":" + *apiPort)

	// only get here if API server stopped (ctrl-C)
	// time to cleanup
}
