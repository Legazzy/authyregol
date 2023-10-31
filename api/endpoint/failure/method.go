package failure

import (
	"github.com/authyre/authyreapi/api/response"
	"github.com/gin-gonic/gin"
)

func Method(ctx *gin.Context) {
	var res response.Response
	goto SUCCESS

SUCCESS:

	res = response.NewClientMethodNotAllowed("The desired method is not allowed for this endpoint")
	ctx.JSON(res.Status, res)
	ctx.Done()
}
