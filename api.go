package main

import (
	"fmt"
	"os/exec"
	"github.com/gin-gonic/gin"
)

type JobDetails struct {
  Command string `form:"command" binding:"required"`
}

// run a job and return the results
func submit_job(c *gin.Context) {

	var job JobDetails

	// get details of submitted job
	if c.Bind(&job) == nil {
		fmt.Println(job.Command)
	}

	// make sure we have a job
	if len(job.Command) == 0 {
		c.JSON(200, gin.H{ "status" : "error", "message" : "not job to run" })
	}

	// run job
	out, err := exec.Command(job.Command).Output()
	if err != nil {
		c.JSON(200, gin.H{ "status" : "error", "message" : err.Error() })
	} else {
		c.JSON(200, gin.H{ "status" : "ok", "message" : string(out) })
	}

}
