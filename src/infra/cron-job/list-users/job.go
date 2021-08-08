package listusers_job

import (
	cronjob "go-api/src/infra/cron-job"
	"go-api/src/main/container"

	"github.com/robfig/cron"
)

type listUserJob struct {
	Cron      *cron.Cron
	container container.Container
}

func New(c container.Container) cronjob.CronJob {
	cron := cron.New()

	return &listUserJob{
		Cron:      cron,
		container: c,
	}
}

func (listUserJob *listUserJob) Start() {
	//
	listUserJob.Cron.Start()
}

func (listUserJob *listUserJob) Stop() {
	//
	listUserJob.Cron.Stop()
}
