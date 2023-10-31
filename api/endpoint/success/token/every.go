package token

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

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionPersonalTokens().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto TARGET
	}

	res = response.NewClientForbidden("You do not have permission to access this resource")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	var tar []token.Token
	if fetch.TokensByIdentifierUser(&tar, usr.Identifier.User) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while fetching tokens")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessOK("All tokens fetched successfully")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
