package population

import (
	"github.com/authyre/authyre-api/pkg/database/request/fetch"
	"github.com/authyre/authyre-api/pkg/database/request/insert"
	"github.com/authyre/authyre-api/pkg/object/user"
	"github.com/authyre/authyre-api/pkg/setup/standard"
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
