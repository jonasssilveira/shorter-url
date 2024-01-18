package cron

import (
	"gopkg.in/robfig/cron.v2"
)

type ExecutionCron interface {
	AddFunc(spec string, cmd func()) (cron.EntryID, error)
	Start()
}
