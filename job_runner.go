package main

import (
	"time"
)

var Jobs []JobDetails

func JobRunner() {
	// every minute 
	// go through list of jobs
	for _, job := range Jobs {
		// is one due to be run
		if job.IsReadyToRun() {
			go job.Run()
		}

		// is it already running
		
	}

	time.Sleep(60)
}


