package token

import (
	"github.com/Authyre/authyreapi/api/middleware/authentication"
	"github.com/Authyre/authyreapi/api/middleware/connection"
	"github.com/gin-gonic/gin"
)

func Attach(eng *gin.Engine) {
	grp := eng.Group("/token")

	cre := grp.Group("")
	tok := grp.Group("")

	cre.Use(authentication.HandleLimits)
	cre.Use(authentication.HandleBasic)

	tok.Use(authentication.HandleLimits)
	tok.Use(authentication.HandleBearer)
	tok.Use(connection.HandleEndpoint)

	cre.POST("", Insert)
	tok.GET("", Every)
	tok.GET("/:credential_access", Fetch)
	tok.DELETE("/:credential_access", Remove)
}
