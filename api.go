package main

import (
	//	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)


// run a job and return the results
func submit_job(c *gin.Context) {

	var job JobDetails

	// get details of submitted job
	if c.Bind(&job) == nil {
		fmt.Println(job.Command)
	}

	// make sure we have a valid job
	if len(job.Command) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "not job to run"})
	}


	// add job to queue

	out, err := exec.Command(job.Command).Output()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": string(out)})
	}

}

// get a list of job names
func list_jobs(c *gin.Context) {

	var jobNames []string

	// go through job list
	for _, job := range Jobs {
		// job name
		jobNames = append( jobNames, job.Name)
	}
	b, err := json.Marshal( jobNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	// return list
	c.JSON( http.StatusOK, gin.H{"status" : "ok", "jobs" : string(b)})
}

// given a job name, return details
func job_details(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "")
}

// delete given job
func delete_job(c *gin.Context) {
	// get name of job to delete
	// find in list
	// delete
	c.JSON(http.StatusNotImplemented, "")
}
