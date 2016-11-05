package main

import (
	"github.com/gin-gonic/gin"
	"os/exec"
)

func submit_job(c *gin.Context) {
	out, err := exec.Command("date").Output()
	if err != nil {

		c.JSON(200, "")
	} else {
		// return date
		c.JSON(200, out)
	}

}
