package user

import (
	"github.com/authyre/authyre-api/api/response"
	"github.com/authyre/authyre-api/api/transfer"
	"github.com/authyre/authyre-api/pkg/database/request/fetch"
	"github.com/authyre/authyre-api/pkg/database/request/modify"
	"github.com/authyre/authyre-api/pkg/object/permission"
	"github.com/authyre/authyre-api/pkg/object/token"
	"github.com/authyre/authyre-api/pkg/object/user"
	"github.com/authyre/authyre-api/pkg/setup/standard"
	"github.com/authyre/authyre-api/pkg/setup/tool"
	"github.com/gin-gonic/gin"
)

func Modify(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)
	usr := ctx.MustGet("user").(user.User)

	var res response.Response
	goto TARGET

TARGET:

	var tar user.User
	if fetch.UserByCredentialUsername(&tar, ctx.Param("credential_username")) == nil {
		goto PERMISSION
	}

	res = response.NewClientNotFound("The user was not found")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServiceUsers().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto REQUEST
	}

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionPersonalChanges().Details.Keyword, tok.Identifier.Permissions) == nil && usr.Identifier.User == tar.Identifier.User {
		goto REQUEST
	}

	res = response.NewClientForbidden("You do not have permission to modify this user")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

REQUEST:

	var req transfer.User

	req.Address = tar.Address
	req.Contact = tar.Contact
	req.Personal = tar.Personal

	if ctx.ShouldBindJSON(req) == nil {
		tar.Address = req.Address
		tar.Contact = req.Contact
		tar.Personal = req.Personal

		if tar.Credential.Password != req.Credential.Password {
			tar.Credential.Password = tool.NewHash(req.Credential.Password + tar.Credential.Prefixes)
		}

		goto PROCESSING
	}

	res = response.NewClientBadRequest("The provided modification data was invalid")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PROCESSING:

	if modify.UserByIdentifierUser(&tar, tar.Identifier.User) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An internal server error occurred while modifying the user")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessNoContent("The user was successfully modified")
	ctx.JSON(res.Status, res)
	ctx.Done()
}
