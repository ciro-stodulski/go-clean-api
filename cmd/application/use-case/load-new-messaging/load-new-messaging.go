package loadnewmessagingusecase

import (
	messagingentity "go-clean-api/cmd/domain/entity/messaging"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/presentation/http/controller"
	"log"
)

type loadNewMessagingUseCase struct {
	channelEvents map[string]controller.ChannelManager[messagingentity.MessagingEntity]
}

func New(channelEvents map[string]controller.ChannelManager[messagingentity.MessagingEntity]) usecase.UseCase[string, messagingentity.MessagingEntity] {
	return &loadNewMessagingUseCase{channelEvents}
}

func (vr *loadNewMessagingUseCase) Perform(input string) (messagingentity.MessagingEntity, error) {

	channel, exists := vr.channelEvents[input]
	if !exists {

		channel.Channel = make(chan messagingentity.MessagingEntity)

		channel.IsOpen = true

		vr.channelEvents[input] = channel

		log.Default().Println("Create channel for subject_id:", input)
	} else {
		log.Default().Println("Channel already exists with subject_id:", input)

	}

	return <-vr.channelEvents[input].Channel, nil
}
