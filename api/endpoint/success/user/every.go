package user

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/object/permission"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/object/user"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Every(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto PERMISSION

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServiceUsers().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto TARGET
	}

	res = response.NewClientForbidden("You do not have permission to access this resource")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	var tar []user.User
	if fetch.Users(&tar) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while fetching users")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessOK("All users have been fetched")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
