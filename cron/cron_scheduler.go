package cron

import (
	"github.com/robfig/cron/v3"
)

var (
	Scheduler = cron.New()
)

func init() {
	Scheduler.AddJob("00 07 * * *", &Hello{})
}
