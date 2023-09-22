package usecase

type (
	IUseCase[I any, O any] interface {
		Perform(I) (O, error)
	}
)
