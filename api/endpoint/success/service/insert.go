package service

import (
	"github.com/Authyre/authyreapi/api/middleware/authentication"
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/api/transfer"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/insert"
	"github.com/Authyre/authyreapi/pkg/object/permission"
	"github.com/Authyre/authyreapi/pkg/object/service"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
	"regexp"
)

func Insert(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto PERMISSION

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServiceServices().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto REQUEST
	}

	res = response.NewClientForbidden("You do not have permission to remove services")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

REQUEST:

	req := transfer.Service{}
	if ctx.ShouldBindJSON(&req) == nil {
		goto VALIDATION
	}

	res = response.NewClientBadRequest("The request body could not be parsed")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

VALIDATION:

	if mat, _ := regexp.MatchString("\\w{5,20}", req.Name); mat {
		goto CONFLICT
	}

	res = response.NewClientBadRequest("The name must be between 5 and 20 characters and contain word characters")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

CONFLICT:

	if fetch.ServiceByDetailsName(&service.Service{}, req.Name) != nil {
		goto TARGET
	}

	res = response.NewClientConflict("The service already exists")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	tar := service.NewService()
	tar.Details.Address = req.Address
	tar.Details.Author = req.Author
	tar.Details.Description = req.Details
	tar.Details.Name = req.Name
	tar.Details.Version = req.Version

	if insert.Service(&tar) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while creating the service")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	authentication.Services[tar.Details.Address] = &tar

	res = response.NewSuccessCreated("The service was successfully created")
	res.Return = req
	ctx.JSON(res.Status, res)
	ctx.Done()
}
