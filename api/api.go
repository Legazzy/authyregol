package api

import (
	"fmt"
	"github.com/authyre/authyreapi/api/endpoint/failure"
	"github.com/authyre/authyreapi/api/endpoint/success/permission"
	"github.com/authyre/authyreapi/api/endpoint/success/service"
	"github.com/authyre/authyreapi/api/endpoint/success/token"
	"github.com/authyre/authyreapi/api/endpoint/success/user"
	"github.com/gin-gonic/gin"
)

var Hostname = ""
var InetPort = "8080"

func Attach() {
	eng := gin.Default()

	eng.NoRoute(failure.Missing)
	eng.NoMethod(failure.Method)

	adr := fmt.Sprintf("%s:%s", Hostname, InetPort)

	permission.Attach(eng)
	service.Attach(eng)
	token.Attach(eng)
	user.Attach(eng)

	if err := eng.Run(adr); err != nil {
		panic(err)
	}
}
