package work

import (
	cronjob "go-api/src/infra/cron-job"
	listusers_job "go-api/src/infra/cron-job/list-users"
	"go-api/src/main/container"
	"log"
	"strconv"
)

type work struct {
	Jobs  []cronjob.CronJob
	Count int
}

func New(c *container.Container) Work {
	jobs := []cronjob.CronJob{listusers_job.New(*c)}

	return &work{
		Jobs:  jobs,
		Count: len(jobs),
	}
}

func (work *work) StartCrons() {
	for _, job := range work.Jobs {
		job.Start()
	}

	log.Default().Print("Jobs started:" + strconv.Itoa(work.Count))
}

func (work *work) StopCrons() {
	for _, job := range work.Jobs {
		job.Stop()
	}

	log.Default().Print("Jobs stopped:" + strconv.Itoa(work.Count))
}
