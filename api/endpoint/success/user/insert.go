package user

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/api/transfer"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/insert"
	"github.com/Authyre/authyreapi/pkg/object/permission"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/object/user"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
	"github.com/Authyre/authyreapi/pkg/setup/tool"
	"github.com/gin-gonic/gin"
	"regexp"
)

func Insert(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto PERMISSION

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServiceUsers().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto REQUEST
	}

	res = response.NewClientForbidden("You do not have permission to create users")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

REQUEST:

	req := transfer.User{}
	if ctx.ShouldBindJSON(&req) == nil {
		goto VALIDATION
	}

	res = response.NewClientBadRequest("The request body was not valid")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

VALIDATION:

	if mat, _ := regexp.MatchString("^[\\w]{5,20}$", req.Credential.Username); mat {
		goto CONFLICT
	}

	res = response.NewClientBadRequest("The username must be between 5 and 20 characters and contain word characters")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

CONFLICT:

	if fetch.UserByCredentialUsername(&user.User{}, req.Credential.Username) != nil {
		goto TARGET
	}

	res = response.NewClientConflict("The username is already in use")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	tar := user.NewUser()
	tar.Address = req.Address
	tar.Contact = req.Contact
	tar.Personal = req.Personal
	tar.Credential.Username = req.Credential.Username
	tar.Credential.Password = tool.NewHash(req.Credential.Password + tar.Credential.Prefixes)

	if insert.User(&tar) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while creating the user")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessCreated("The user was created successfully")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()

}
