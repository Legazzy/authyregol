package user

import (
	"github.com/google/uuid"
)

func NewUser() User {
	usr := User{}

	usr.Credential.Prefixes = uuid.New().String()
	usr.Identifier.User = uuid.New().String()

	return usr
}
