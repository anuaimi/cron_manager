package main

// cron manager

import (
	"flag"
	"os"
	"github.com/gin-gonic/gin"
	log "github.com/Sirupsen/logrus"
)

func main() {

	// init logging
	//log.SetLevel(log.WarnLevel)
	log.SetLevel(log.InfoLevel)
	log.Info("cron_manager starting")

	// read parameters
	port := os.Getenv("CM_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	// get input parameters
	apiKey := flag.String("api_key", os.Getenv("CM_APIKEY"), "api key to authenticate")
	apiPort := flag.String("port", port, "port for API")
	flag.Parse()

	// load previous jobs from config file

	// start job manager
	

	// setup api routes
	r := gin.Default()
	authorized := r.Group("/")
	if len(*apiKey) > 0 {
		authorized = r.Group("/", gin.BasicAuth(gin.Accounts{"apiKey": *apiKey}))
	}

	authorized.GET("/jobs", list_jobs)
	authorized.GET("/job", job_details)
	authorized.POST("/job", submit_job)
	authorized.DELETE("/job", delete_job)

	r.Run(":" + *apiPort)

	// only get here if API server stopped (ctrl-C)
	// time to cleanup
}
