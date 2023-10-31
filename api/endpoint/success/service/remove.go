package service

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/remove"
	"github.com/Authyre/authyreapi/pkg/object/permission"
	"github.com/Authyre/authyreapi/pkg/object/service"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
	"github.com/gin-gonic/gin"
)

func Remove(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto PERMISSION

PERMISSION:

	if fetch.PermissionByDetailKeywordAndIdentifierPermissions(&permission.Permission{}, standard.GetPermissionServiceServices().Details.Keyword, tok.Identifier.Permissions) == nil {
		goto TARGET
	}

	res = response.NewClientForbidden("You do not have permission to remove services")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

TARGET:

	var tar service.Service
	if fetch.ServiceByDescriptionName(&tar, ctx.Param("description_name")) == nil {
		goto PROCESSING
	}

	res = response.NewClientNotFound("The service you are trying to remove does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PROCESSING:

	if remove.ServiceByIdentifierService(tar.Identifier.Service) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("An error occurred while removing the service")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessNoContent("The service was successfully removed")
	ctx.JSON(res.Status, res)
	ctx.Done()
}
