package population

import (
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/database/request/insert"
	"github.com/authyre/authyreapi/pkg/object/service"
	"github.com/authyre/authyreapi/pkg/setup/standard"
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
