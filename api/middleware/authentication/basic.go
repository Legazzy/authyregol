package authentication

import (
	"github.com/authyre/authyre-api/api/response"
	"github.com/authyre/authyre-api/pkg/database/request/fetch"
	"github.com/authyre/authyre-api/pkg/object/user"
	"github.com/authyre/authyre-api/pkg/setup/tool"
	"github.com/gin-gonic/gin"
)

func HandleBasic(ctx *gin.Context) {
	nme, pwd, _ := ctx.Request.BasicAuth()

	var res response.Response
	goto TARGET

TARGET:

	var usr user.User
	if fetch.UserByCredentialUsername(&usr, nme) == nil {
		goto VALIDATION
	}

	res = response.NewClientUnauthorized("The provided username does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

VALIDATION:

	if usr.Credential.Password == tool.NewHash(pwd+usr.Credential.Prefixes) {
		goto NEXT
	}

	res = response.NewClientUnauthorized("The provided password is incorrect")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

NEXT:

	ctx.Set("user", usr)
	ctx.Next()
}
