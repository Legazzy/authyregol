package permission

import (
	"github.com/google/uuid"
)

func NewPermission() Permission {
	per := Permission{}

	per.Identifier.Permission = uuid.New().String()

	return per
}
