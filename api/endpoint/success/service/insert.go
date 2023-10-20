package service

import (
	"github.com/authyre/authyre-api/api/response"
	"github.com/authyre/authyre-api/api/transfer"
	"github.com/authyre/authyre-api/pkg/database/request/fetch"
	"github.com/authyre/authyre-api/pkg/database/request/insert"
	"github.com/authyre/authyre-api/pkg/object/permission"
	"github.com/authyre/authyre-api/pkg/object/service"
	"github.com/authyre/authyre-api/pkg/object/token"
	"github.com/authyre/authyre-api/pkg/setup/standard"
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

	reg := "\\w{5,20}"
	if mat, _ := regexp.MatchString(reg, req.Name); mat {
		goto CONFLICT
	}

	res = response.NewClientBadRequest("The name must be between 5 and 20 characters and contain word characters")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

CONFLICT:

	if fetch.ServiceByDescriptionName(&service.Service{}, req.Name) != nil {
		goto TARGET
	}

	res = response.NewClientConflict("The service already exists")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	tar := service.NewService()
	tar.Description.Author = req.Author
	tar.Description.Details = req.Details
	tar.Description.Name = req.Name
	tar.Description.Version = req.Version

	if insert.Service(&tar) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while creating the service")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessCreated("The service was successfully created")
	res.Return = req
	ctx.JSON(res.Status, res)
	ctx.Done()
}
