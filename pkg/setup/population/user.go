package population

import (
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/insert"
	"github.com/Authyre/authyreapi/pkg/object/user"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
)

var RepopulateUsers = true

func InsertUsers() {
	if !RepopulateUsers {
		return
	}

	usr := standard.GetUser()
	if fetch.UserByIdentifierUser(&user.User{}, usr.Identifier.User) != nil {
		_ = insert.User(&usr)
	}
}
