package transfer

import "github.com/Authyre/authyreapi/pkg/object/user"

type User struct {
	Address  user.Address
	Contact  user.Contact
	Personal user.Personal

	Credential struct {
		Password string
		Username string
	}
}
