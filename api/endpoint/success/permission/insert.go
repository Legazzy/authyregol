package permission

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/api/transfer"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/insert"
	"github.com/Authyre/authyreapi/pkg/object/permission"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Insert(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto REQUEST

REQUEST:

	var req transfer.Permission
	if ctx.ShouldBindJSON(&req) == nil {
		goto PERMISSION
	}

	res = response.NewClientBadRequest("The request body was not formatted correctly")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServicePermissions().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto FORBIDDEN
	}

	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&permission.Permission{}, req.Keyword, ctx.Param("identifier_service"), ctx.Param("identifier_user")) == nil {
		goto FORBIDDEN
	}

	goto TARGET

FORBIDDEN:

	res = response.NewClientForbidden("You don't have permission to create this permission")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	tar := permission.NewPermission()
	tar.Identifier.User = ctx.Param("identifier_user")
	tar.Identifier.Service = ctx.Param("identifier_service")
	tar.Details.Description = req.Description
	tar.Details.Keyword = req.Keyword

	if insert.Permission(&tar) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("The server encountered an error while creating the permission")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessCreated("Created permission")
	ctx.JSON(res.Status, res)
	ctx.Done()
}
