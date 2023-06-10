package service

import (
	"go-clean-api/cmd/domain/dto"
)

type (
	NotificationService interface {
		SendNotify(dto dto.Event) error
		CheckNotify(msg string) error
		SaveNotify(dto.Event) string
		FindById(id string) (*dto.Event, error)
	}
)
