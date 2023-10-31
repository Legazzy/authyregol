package standard

import "github.com/Authyre/authyreapi/pkg/object/service"

func GetService() service.Service {
	ser := service.NewService()

	ser.Description.Author = "Otto Rohenkohl"
	ser.Description.Details = "A service for authenticating users"
	ser.Description.Name = "authyre"
	ser.Description.Version = "0.0.1"
	ser.Identifier.Service = "00000000-0000-0000-0000-000000000000"

	return ser
}
