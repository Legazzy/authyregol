package authentication

import (
	"github.com/Authyre/authyreapi/api/response"
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/remove"
	"github.com/Authyre/authyreapi/pkg/object/token"
	"github.com/Authyre/authyreapi/pkg/object/user"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func HandleBearer(ctx *gin.Context) {
	acc := strings.Split(ctx.GetHeader("Authorization")+" ", " ")[1]

	var res response.Response
	goto TARGET

TARGET:

	var tok token.Token
	if fetch.TokenByCredentialAccess(&tok, acc) == nil {
		goto VALIDATION
	}

	res = response.NewClientUnauthorized("The provided token is invalid or has expired")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

VALIDATION:

	if time.Now().Before(time.Unix(tok.Timestamp.Expiration, 0)) {
		goto PROCESSING
	}

	go func() { _ = remove.TokenByIdentifierToken(tok.Identifier.Token) }()

	res = response.NewClientUnauthorized("The provided token is invalid or has expired")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

PROCESSING:

	var usr user.User
	if fetch.UserByIdentifierUser(&usr, tok.Identifier.User) == nil {
		goto NEXT
	}

	res = response.NewClientUnauthorized("The provided token is invalid or has expired")
	ctx.JSON(res.Status, res)
	ctx.Abort()

	return

NEXT:

	ctx.Set("token", tok)
	ctx.Set("user", usr)
	ctx.Next()
}
