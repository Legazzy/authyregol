package user

import (
	"github.com/Authyre/authyreapi/api/middleware/authentication"
	"github.com/Authyre/authyreapi/api/middleware/connection"
	"github.com/gin-gonic/gin"
)

func Attach(eng *gin.Engine) {
	grp := eng.Group("/user")

	grp.Use(authentication.HandleLimits)
	grp.Use(authentication.HandleBearer)
	grp.Use(connection.HandleEndpoint)

	grp.GET("", Every)
	grp.GET("/:credential_username", Fetch)
	grp.POST("", Insert)
	grp.PATCH("/:credential_username", Modify)
	grp.DELETE("/:credential_username", Remove)
}
