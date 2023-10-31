package service

import (
	"github.com/authyre/authyreapi/api/response"
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/object/service"
	"github.com/gin-gonic/gin"
)

func Fetch(ctx *gin.Context) {
	var res response.Response
	goto TARGET

TARGET:

	var tar service.Service
	if fetch.ServiceByDescriptionName(&tar, ctx.Param("description_name")) == nil {
		goto SUCCESS
	}

	res = response.NewClientNotFound("The service you are looking for does not exist")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessOK("The service was found successfully")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
