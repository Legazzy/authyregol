package standard

import "github.com/Authyre/authyreapi/pkg/object/service"

func GetService() service.Service {
	ser := service.NewService()

	ser.Details.Author = "Otto Rohenkohl"
	ser.Details.Description = "A service for authenticating users"
	ser.Details.Name = "authyre"
	ser.Details.Version = "0.0.1"
	ser.Identifier.Service = "00000000-0000-0000-0000-000000000000"

	return ser
}
