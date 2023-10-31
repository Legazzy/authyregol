package population

import (
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/database/request/insert"
	"github.com/authyre/authyreapi/pkg/object/user"
	"github.com/authyre/authyreapi/pkg/setup/standard"
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
