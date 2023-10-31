package user

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/remove"
	"github.com/Authyre/authyreapi/pkg/object/permission"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/object/user"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Remove(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)
	usr := ctx.MustGet("user").(user.User)

	var res response.Response
	goto TARGET

TARGET:

	var tar user.User
	if fetch.UserByCredentialUsername(&tar, ctx.Param("credential_username")) == nil {
		goto PERMISSION
	}

	res = response.NewClientNotFound("The user does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServiceUsers().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto PROCESSING
	}

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionPersonalChanges().Details.Keyword, tok.Identifier.Permissions) == nil && usr.Identifier.User == tar.Identifier.User {
		goto PROCESSING
	}

	res = response.NewClientForbidden("You do not have permission to remove this user")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PROCESSING:

	if remove.UserByIdentifierUser(tar.Identifier.User) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while removing the user")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessNoContent("The user was removed successfully")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
