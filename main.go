package main

// cron manager

import (
	"flag"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {

	// read parameters
	port := flag.String("port",os.Getenv("CM_PORT"),"port for API")
	flag.Parse()

	if len(*port) == 0 {
		*port = "8080"
	}

	// setup api routes
	r := gin.Default()
	r.POST("/job", submit_job)
	r.Run(":" + *port)

	// only get here if API server stopped (ctrl-C)
	// time to cleanup
}
