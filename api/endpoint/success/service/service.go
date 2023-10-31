package service

import (
	"github.com/authyre/authyreapi/api/middleware/authentication"
	"github.com/authyre/authyreapi/api/middleware/connection"
	"github.com/gin-gonic/gin"
)

func Attach(eng *gin.Engine) {
	grp := eng.Group("/service")

	grp.GET("", Every)
	grp.GET("/:description_name", Fetch)

	tok := grp.Group("")

	grp.Use(authentication.HandleLimits)
	tok.Use(authentication.HandleBearer)
	tok.Use(connection.HandleEndpoint)

	tok.POST("", Insert)
	tok.DELETE("/:description_name", Remove)
}
