package authentication

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/object/request"
	"github.com/Authyre/authyreapi/pkg/object/service"
	"github.com/gin-gonic/gin"
	"time"
)

var Requests = make(map[string]*request.Request)
var Services = make(map[string]*service.Service)

func HandleLimits(ctx *gin.Context) {
	req := Requests[ctx.ClientIP()]

	var res response.Response
	goto WILDCARD

WILDCARD:

	if Services[ctx.ClientIP()] != nil {
		goto NEXT
	}

	goto TARGET

TARGET:

	if req != nil {
		goto VALIDATION
	}

	Requests[ctx.ClientIP()] = request.NewRequest()
	ctx.Next()

	return

VALIDATION:

	req.Amount++
	req.Amount -= (time.Now().Unix() - req.Last) / req.Cooldown
	req.Last = time.Now().Unix()

	if req.Amount <= req.Limit {
		goto NEXT
	}

	res = response.NewClientTooManyRequests("The request limit is exceeded")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

NEXT:

	ctx.Next()
}
