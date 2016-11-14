package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateSampleJobs(t *testing.T) {

	// every 15 minutes
	sched1 := &JobDetails{ Name: "job1", Minutes: []int{00,15,30,45}, Command: "pwd"} 
	// every hour 
	sched1 = &JobDetails{ Name: "job2", Hours: []int{00}, Command: "pwd"} 
	// every day at nooon
	sched1  = &JobDetails{ Name: "job3", Minutes: []int{00}, Hours: []int{12}, Command: "pwd"}
	fmt.Println( sched1)
}

func TestIsReadyToRun( t *testing.T) {

	currentTime := time.Now()

	job := &JobDetails{ Name: "job1", Minutes: []int{currentTime.Minute()}, Command: ""}
	if !job.IsReadyToRun() {
		t.Error("IsReadyToRun reporting not ready when should be ready")
	}


}

