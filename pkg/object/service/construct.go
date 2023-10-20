package service

import (
	"github.com/google/uuid"
)

func NewService() Service {
	ser := Service{}

	ser.Identifier.Service = uuid.New().String()

	return ser
}
