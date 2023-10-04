package eventscontroller

import (
	messagingentity "go-clean-api/cmd/domain/entity/messaging"
	"go-clean-api/cmd/domain/exception"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/presentation/http/controller"
	"log"
)

type eventsController struct {
	loadNewMessagingUseCase usecase.UseCase[string, messagingentity.MessagingEntity]
}

func New(loadNewMessagingUseCase usecase.UseCase[string, messagingentity.MessagingEntity]) controller.Controller {
	return &eventsController{
		loadNewMessagingUseCase,
	}
}

// LoadRoute implements controller.Controller.
func (ec *eventsController) LoadRoute() controller.CreateRoute {
	return controller.CreateRoute{
		PathRoot:              "/v1/events",
		Method:                "get",
		Path:                  "/:subject_id",
		IsServerSentEvents:    true,
		TimeSecondsSentEvents: 2,
	}
}

// Handle implements controller.Controller.
func (ec *eventsController) Handle(req controller.HttpRequest) (*controller.HttpResponse, error) {
	result, err := ec.loadNewMessagingUseCase.Perform(req.Params.Get("subject_id"))

	return &controller.HttpResponse{
		Data:   result,
		Status: 200,
	}, err
}

// HandleError implements controller.Controller.
func (ec *eventsController) HandleError(appErr *exception.ApplicationException) *controller.HttpResponseError {
	log.Printf("[eventsController]{HandleError}: error internal %v", appErr)

	return nil
}
