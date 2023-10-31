package permission

import (
	"github.com/authyre/authyreapi/api/response"
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/object/permission"
	"github.com/authyre/authyreapi/pkg/object/token"
	"github.com/authyre/authyreapi/pkg/object/user"
	"github.com/authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Every(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)
	usr := ctx.MustGet("user").(user.User)

	var res response.Response
	goto PERMISSION

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServicePermissions().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto TARGET
	}

	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&permission.Permission{}, standard.GetPermissionPersonalPermissions().Details.Keyword, ctx.Param("identifier_service"), ctx.Param("identifier_user")) == nil && ctx.Param("identifier_user") == usr.Identifier.User {
		goto TARGET
	}

	res = response.NewClientForbidden("You don't have permission to view those permissions")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	var tar []permission.Permission
	if fetch.PermissionsByIdentifierServiceAndIdentifierUser(&tar, ctx.Param("identifier_service"), ctx.Param("identifier_user")) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while retrieving permissions")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessCreated("All permissions where retrieved successfully")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
