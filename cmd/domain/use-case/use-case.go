package usecase

type (
	UseCase[I any, O any] interface {
		Perform(I) (O, error)
	}
)
