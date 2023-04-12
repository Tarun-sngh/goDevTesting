package cron

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

func TestCronJob() *cron.Cron {
	cronJob := cron.New()
	_, err := cronJob.AddFunc("* * * * *", func() {
		fmt.Println("Current time is -------------------------------: ", time.Now())
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		panic(err)
	}
	return cronJob
}

func StartTestCronJob() {
	var cron *cron.Cron = TestCronJob()
	cron.Start()
}
