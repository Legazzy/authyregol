package connection

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/database/request/insert"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/gin-gonic/gin"
)

func HandleEndpoint(ctx *gin.Context) {
	tok := ctx.MustGet("token").(token.Token)

	var res response.Response
	goto PROCESSING

PROCESSING:

	tok.Endpoint.Address = ctx.ClientIP()
	tok.Endpoint.Device = ctx.GetHeader("User-Agent")

	if insert.Token(&tok) == nil {
		goto NEXT
	}

	res = response.NewServerInternalServerError("Failed to update endpoint information")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

NEXT:

	ctx.Next()
}
