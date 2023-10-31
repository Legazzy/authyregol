package service

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/object/service"
	"github.com/gin-gonic/gin"
)

func Every(ctx *gin.Context) {
	var res response.Response
	goto TARGET

TARGET:

	var tar []service.Service
	if fetch.Services(&tar) == nil {
		goto SUCCESS
	}

	res = response.NewServerInternalServerError("The services could not be fetched")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	res = response.NewSuccessOK("The services were fetched successfully")
	res.Return = tar
	ctx.JSON(res.Status, res)
	ctx.Done()
}
