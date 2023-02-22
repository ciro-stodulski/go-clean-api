package listusersjob

import (
	"go-clean-api/cmd/main/container"
	cronjob "go-clean-api/cmd/presentation/cron-job"
	"log"

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

func (luj *listUserJob) Start() {
	luj.Cron.AddFunc("1 * * * *", func() {
		log.Default().Print("### job ListUsers started ###")
		luj.container.ListUsersUseCase.ListUsers()
		log.Default().Print("### job ListUsers finishid ###")
	})

	luj.Cron.Start()
}

func (luj *listUserJob) Stop() {
	luj.Cron.Stop()
}
