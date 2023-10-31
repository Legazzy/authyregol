package standard

import (
	"github.com/authyre/authyreapi/pkg/object/user"
	"github.com/authyre/authyreapi/pkg/setup/tool"
)

func GetUser() user.User {
	usr := user.NewUser()
	hsh := tool.NewHash("authyre" + usr.Credential.Prefixes)

	usr.Credential.Password = hsh
	usr.Credential.Username = "authyre"
	usr.Identifier.User = "00000000-0000-0000-0000-000000000000"

	return usr
}
