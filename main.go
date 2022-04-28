package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func main() {
	sleepInterval := 500
	if os.Getenv("SLEEP_INTERVAL") != "" {
		interval, _ := strconv.Atoi(os.Getenv("SLEEP_INTERVAL"))
		if interval != 0 {
			sleepInterval = interval
		}
	}
	for {
		logrus.Info(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Millisecond * time.Duration(sleepInterval))
	}
}
