package permission

import (
	"github.com/Authyre/authyreapi/api/middleware/authentication"
	"github.com/Authyre/authyreapi/api/middleware/connection"
	"github.com/gin-gonic/gin"
)

func Attach(eng *gin.Engine) {
	grp := eng.Group("/permission")

	grp.Use(authentication.HandleLimits)
	grp.Use(authentication.HandleBearer)
	grp.Use(connection.HandleEndpoint)

	grp.GET("/:identifier_service/:identifier_user", Every)
	grp.POST("/:identifier_service/:identifier_user", Insert)
	grp.DELETE("/:permission_identifier", Remove)
}
