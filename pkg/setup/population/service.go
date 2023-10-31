package population

import (
	"github.com/Authyre/authyreapi/pkg/database/request/fetch"
	"github.com/Authyre/authyreapi/pkg/database/request/insert"
	"github.com/Authyre/authyreapi/pkg/object/service"
	"github.com/Authyre/authyreapi/pkg/setup/standard"
)

var RepopulateServices = true

func InsertServices() {
	if !RepopulateServices {
		return
	}

	ser := standard.GetService()
	if fetch.ServiceByIdentifierService(&service.Service{}, ser.Identifier.Service) != nil {
		_ = insert.Service(&ser)
	}
}
