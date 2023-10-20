package user

import (
	"github.com/authyre/authyre-api/api/response"
	"github.com/authyre/authyre-api/pkg/database/request/fetch"
	"github.com/authyre/authyre-api/pkg/object/permission"
	"github.com/authyre/authyre-api/pkg/object/token"
	"github.com/authyre/authyre-api/pkg/object/user"
	"github.com/authyre/authyre-api/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Fetch(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)
	usr := ctx.MustGet("user").(user.User)

	var res response.Response
	goto TARGET

TARGET:

	var tar user.User
	if fetch.UserByCredentialUsername(&tar, ctx.Param("credential_username")) == nil {
		goto PERMISSION
	}

	res = response.NewClientNotFound("The desired user was not found")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PERMISSION:

	if (fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServiceUsers().Details.Keyword, tok.Identifier.Permissions)) == nil {
		goto SUCCESS
	}

	if (fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionPersonalInfos().Details.Keyword, tok.Identifier.Permissions) == nil && usr.Identifier.User == tar.Identifier.User) {
		goto SUCCESS
	}

	res = response.NewClientForbidden("You do not have permission to view this user")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessOK("The desired user was found")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
