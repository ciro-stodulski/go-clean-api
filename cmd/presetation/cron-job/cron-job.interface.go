package cronjob

type CronJob interface {
	Start()
	Stop()
}
