package listusersjob

import (
	usecase "go-clean-api/cmd/domain/use-case"
	cronjob "go-clean-api/cmd/presentation/cron-job"
	"log"

	"github.com/robfig/cron"
)

type listUserJob struct {
	Cron             *cron.Cron
	listUsersUseCase usecase.UseCase[any, any]
}

func New(listUsersUseCase usecase.UseCase[any, any]) cronjob.CronJob {
	cron := cron.New()

	return &listUserJob{
		Cron:             cron,
		listUsersUseCase: listUsersUseCase,
	}
}

func (luj *listUserJob) Start() {
	luj.Cron.AddFunc("1 * * * *", func() {
		log.Default().Print("### job ListUsers started ###")
		luj.listUsersUseCase.Perform(nil)
		log.Default().Print("### job ListUsers finishid ###")
	})

	luj.Cron.Start()
}

func (luj *listUserJob) Stop() {
	luj.Cron.Stop()
}
