package work

import (
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/main/modules"
	cronjob "go-clean-api/cmd/presentation/cron-job"
	listusers_job "go-clean-api/cmd/presentation/cron-job/list-users"
	"log"
	"strconv"
)

type work struct {
	Jobs  []cronjob.CronJob
	Count int
}

func New(c *container.Container) modules.Module {
	jobs := []cronjob.CronJob{listusers_job.New(c.ListUsersUseCase)}

	return &work{
		Jobs:  jobs,
		Count: len(jobs),
	}
}

func (work *work) RunGo() bool {
	return false
}

func (work *work) Start() error {
	for _, job := range work.Jobs {
		job.Start()
	}

	log.Default().Print("Jobs started:" + strconv.Itoa(work.Count))

	return nil
}

func (work *work) Stop() {
	for _, job := range work.Jobs {
		job.Stop()
	}

	log.Default().Print("Jobs stopped:" + strconv.Itoa(work.Count))
}
