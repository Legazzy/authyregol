package token

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/remove"
	"github.com/Authyre/authyreapi/pkg/object/permission"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Remove(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto TARGET

TARGET:

	var tar token.Token
	if fetch.TokenByCredentialAccess(&tar, ctx.Param("credential_access")) == nil {
		goto PERMISSION
	}

	res = response.NewClientNotFound("The token does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionPersonalTokens().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto PROCESSING
	}

	if tok.Identifier.Token == tar.Identifier.Token {
		goto PROCESSING
	}

	res = response.NewClientNotFound("The token does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PROCESSING:

	if remove.TokenByIdentifierToken(tok.Identifier.Token) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while deleting the token")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessNoContent("The token has been deleted")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
