package sendeventscontroller

import (
	inputdto "go-clean-api/cmd/domain/dto/input"
	"go-clean-api/cmd/domain/exception"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/presentation/http/controller"
	"log"
)

type sendEventsController struct {
	verifyNewMenssagingUseCase usecase.UseCase[inputdto.MessagingInput, any]
}

func New(loadNewMessagingUseCase usecase.UseCase[inputdto.MessagingInput, any]) controller.Controller {
	return &sendEventsController{
		loadNewMessagingUseCase,
	}
}

// LoadRoute implements controller.Controller.
func (ec *sendEventsController) LoadRoute() controller.CreateRoute {
	return controller.CreateRoute{
		PathRoot: "/v1/events",
		Method:   "post",
		Dto:      &inputdto.MessagingInput{},
	}
}

// Handle implements controller.Controller.
func (ec *sendEventsController) Handle(req controller.HttpRequest) (*controller.HttpResponse[any], error) {

	result, err := ec.verifyNewMenssagingUseCase.Perform(req.Body.(inputdto.MessagingInput))

	return &controller.HttpResponse[any]{
		Data:   result,
		Status: 200,
	}, err
}

// HandleError implements controller.Controller.
func (ec *sendEventsController) HandleError(appErr *exception.ApplicationException) *controller.HttpResponse[controller.HttpError] {
	log.Printf("[eventsController]{HandleError}: error internal %v", appErr)

	return nil
}
