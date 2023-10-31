package token

import (
	"github.com/authyre/authyreapi/api/response"
	"github.com/authyre/authyreapi/api/transfer"
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/database/request/insert"
	"github.com/authyre/authyreapi/pkg/object/permission"
	"github.com/authyre/authyreapi/pkg/object/token"
	"github.com/authyre/authyreapi/pkg/object/user"
	"github.com/gin-gonic/gin"
)

func Insert(ctx *gin.Context) {
	usr := ctx.MustGet("user").(user.User)

	var res response.Response
	goto REQUEST

REQUEST:

	var req transfer.Token
	if ctx.ShouldBindJSON(&req) == nil {
		goto TARGET
	}

	res = response.NewClientBadRequest("The given request is invalid")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	tar := token.NewToken()
	tar.Endpoint = req.Endpoint
	tar.Identifier.User = usr.Identifier.User
	tar.Resource = req.Resource

	for _, one := range req.Permissions {
		var per permission.Permission
		if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&per, one.Keyword, one.Service, usr.Identifier.User) != nil {
			goto FORBIDDEN
		}

		tar.Identifier.Permissions = append(tar.Identifier.Permissions, per.Identifier.Permission)
	}

	goto PROCESSING

FORBIDDEN:

	res = response.NewClientForbidden("You do not have permission to create this token")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PROCESSING:

	if insert.Token(&tar) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while creating the token")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessCreated("The token was successfully created")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
