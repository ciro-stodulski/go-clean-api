package sendnewmessagingusecase

import (
	inputdto "go-clean-api/cmd/domain/dto/input"
	messagingentity "go-clean-api/cmd/domain/entity/messaging"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/presentation/http/controller"
	"time"
)

type sendNewMessagingUseCase struct {
	channelEvents map[string]controller.ChannelManager[messagingentity.MessagingEntity]
}

func New(channelEvents map[string]controller.ChannelManager[messagingentity.MessagingEntity]) usecase.UseCase[inputdto.MessagingInput, interface{}] {
	return &sendNewMessagingUseCase{channelEvents}
}

func (vr *sendNewMessagingUseCase) Perform(input inputdto.MessagingInput) (interface{}, error) {

	subjectID := input.SubjectId

	_, exists := vr.channelEvents[subjectID]
	if !exists {

		return nil, nil
	}

	go func(channel controller.ChannelManager[messagingentity.MessagingEntity]) {
		isOpen := channel.IsOpen

		if isOpen {

			channel.Channel <- messagingentity.MessagingEntity{
				To:        input.To,
				Subject:   input.Subject,
				SubjectId: subjectID,
				Body:      input.Body,
				Date:      time.Now(),
			}
		}

	}(vr.channelEvents[subjectID])

	return nil, nil
}
