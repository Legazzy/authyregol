package authentication

import (
	"github.com/authyre/authyre-api/api/response"
	"github.com/authyre/authyre-api/pkg/object/request"
	"github.com/gin-gonic/gin"
	"time"
)

var requests = make(map[string]*request.Request)

func HandleLimits(ctx *gin.Context) {
	req := requests[ctx.ClientIP()]

	var res response.Response
	goto TARGET

TARGET:

	if req != nil {
		goto VALIDATION
	}

	requests[ctx.ClientIP()] = request.NewRequest()
	ctx.Next()

	return

VALIDATION:

	req.Amount++
	req.Amount -= (time.Now().Unix() - req.Last) / req.Cooldown
	req.Last = time.Now().Unix()

	if req.Amount <= req.Limit {
		goto SUCCESS
	}

	res = response.NewClientTooManyRequests("The request limit is exceeded")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

SUCCESS:

	ctx.Next()
}
