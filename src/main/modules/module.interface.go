package modules

type Module interface {
	Start() error
	Stop()
}
