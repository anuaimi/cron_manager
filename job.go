package main

// define jobs 

import (
	"errors"
	"os/exec"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type JobDetails struct {
	Name	string		// name of the job
	Minutes	[]int
	Hours	[]int
	Dates	[]int
	Months	[]int
	Days	[]int
	Command string		// command to run
}

func createSampleJobs() {

	// every 15 minutes
	sched1 := &JobDetails{ Name: "job1", Minutes: []int{00,15,30,45}, Command: "pwd"} 
	// every hour 
	sched1 = &JobDetails{ Name: "job2", Hours: []int{00}, Command: "pwd"} 
	// every day at nooon
	sched1  = &JobDetails{ Name: "job3", Minutes: []int{00}, Hours: []int{12}, Command: "pwd"}
	fmt.Println( sched1)
}


type ApiJobDetails struct {
	// time to run & interval
	Minute string `form:"minute"`
	Hour   string
	Date   string
	Month  string
	Day    string

	Command string `form:"command" binding:"required"`
}

func (js *JobDetails) IsReadyToRun() bool {

	// get current time
	currentTime := time.Now()

	// see if matches a minute rule
	minuteSpecified := len(js.Minutes) > 0
	if minuteSpecified {
		currentMinute := currentTime.Minute()
		for _, v := range js.Minutes {
			if v == currentMinute {
				return true
			}
		}
	}

	// see if matches an hour rule
	hourSpecified := len(js.Hours) > 0
	if hourSpecified {
		currentHour := currentTime.Hour()
		for _, v := range js.Hours {
			if v == currentHour {	
				return true
			}
		}
	}


	return false
}

// take command and run
func (j *JobDetails) Run() (int, error) {

	// create temp directory to run command in
	tmpDir, err := ioutil.TempDir("", "cron_mngr")
	if err != nil {
		return -1, errors.New("could not setup for job")
	}
	// remove temp directory (and any left over files)
	defer os.RemoveAll(tmpDir)

	// ideally command should not be able to access other directories

	// run command
	cmd := exec.Command("/bin/sh", j.Command)	
	err = cmd.Run()
	if err != nil {
		// opps
		return -1, err
	}
		
	// worked

	return 0, nil
}

