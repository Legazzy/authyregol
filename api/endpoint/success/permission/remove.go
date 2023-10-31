package permission

import (
	"github.com/authyre/authyreapi/api/response"
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/database/request/remove"
	"github.com/authyre/authyreapi/pkg/object/permission"
	"github.com/authyre/authyreapi/pkg/object/token"
	"github.com/authyre/authyreapi/pkg/object/user"
	"github.com/authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Remove(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)
	usr := ctx.MustGet("user").(user.User)

	var res response.Response
	goto TARGET

TARGET:

	var tar permission.Permission
	if fetch.PermissionByIdentifierPermission(&tar, ctx.Param("identifier_permission")) == nil {
		goto PERMISSION
	}

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServicePermissions().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto PROCESSING
	}

	if tar.Identifier.User == usr.Identifier.User {
		goto PROCESSING
	}

	res = response.NewClientForbidden("You do not have permission to delete permissions")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PROCESSING:

	if remove.PermissionByIdentifierPermission(tar.Identifier.Permission) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("The permission could not be deleted because of an error")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessNoContent("The permission was deleted successfully")
	ctx.JSON(res.Status, res)
	ctx.Done()
}
