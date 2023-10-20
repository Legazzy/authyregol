package permission

import (
	"github.com/authyre/authyre-api/api/response"
	"github.com/authyre/authyre-api/api/transfer"
	"github.com/authyre/authyre-api/pkg/database/request/fetch"
	"github.com/authyre/authyre-api/pkg/database/request/insert"
	"github.com/authyre/authyre-api/pkg/object/permission"
	"github.com/authyre/authyre-api/pkg/object/token"
	"github.com/authyre/authyre-api/pkg/setup/standard"
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
