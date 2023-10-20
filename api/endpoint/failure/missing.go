package failure

import (
	"github.com/authyre/authyre-api/api/response"
	"github.com/gin-gonic/gin"
)

func Missing(ctx *gin.Context) {
	var res response.Response
	goto SUCCESS

SUCCESS:

	res = response.NewClientNotFound("The requested resource was not found")
	ctx.JSON(res.Status, res)
	ctx.Done()
}
