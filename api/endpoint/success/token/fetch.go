package token

import (
	"github.com/authyre/authyreapi/api/response"
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/object/permission"
	"github.com/authyre/authyreapi/pkg/object/token"
	"github.com/authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Fetch(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto TARGET

TARGET:

	var tar token.Token
	if fetch.TokenByCredentialAccess(&tar, ctx.Param("credential_access")) == nil {
		goto PERMISSION
	}

	res = response.NewClientNotFound("The provided token does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionPersonalTokens().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto SUCCESS
	}

	if tok.Identifier.Token == tar.Identifier.Token {
		goto SUCCESS
	}

	res = response.NewClientNotFound("The provided token does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessOK("The token was successfully fetched")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
