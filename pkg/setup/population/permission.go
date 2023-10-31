package population

import (
	"github.com/authyre/authyreapi/pkg/database/request/fetch"
	"github.com/authyre/authyreapi/pkg/database/request/insert"
	"github.com/authyre/authyreapi/pkg/setup/standard"
)

var RepopulatePermissions = true

func InsertPermissions() {
	if !RepopulatePermissions {
		return
	}

	ppc := standard.GetPermissionPersonalChanges()
	ppc.Identifier.User = standard.GetUser().Identifier.User
	ppc.Identifier.Service = standard.GetService().Identifier.Service
	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&ppc, ppc.Details.Keyword, ppc.Identifier.Service, ppc.Identifier.User) != nil {
		_ = insert.Permission(&ppc)
	}

	ppi := standard.GetPermissionPersonalInfos()
	ppi.Identifier.User = standard.GetUser().Identifier.User
	ppi.Identifier.Service = standard.GetService().Identifier.Service
	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&ppi, ppi.Details.Keyword, ppi.Identifier.Service, ppi.Identifier.User) != nil {
		_ = insert.Permission(&ppi)
	}

	ppp := standard.GetPermissionPersonalPermissions()
	ppp.Identifier.User = standard.GetUser().Identifier.User
	ppp.Identifier.Service = standard.GetService().Identifier.Service
	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&ppp, ppp.Details.Keyword, ppp.Identifier.Service, ppp.Identifier.User) != nil {
		_ = insert.Permission(&ppp)
	}

	ppt := standard.GetPermissionPersonalTokens()
	ppt.Identifier.User = standard.GetUser().Identifier.User
	ppt.Identifier.Service = standard.GetService().Identifier.Service
	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&ppt, ppt.Details.Keyword, ppt.Identifier.Service, ppt.Identifier.User) != nil {
		_ = insert.Permission(&ppt)
	}

	psp := standard.GetPermissionServicePermissions()
	psp.Identifier.User = standard.GetUser().Identifier.User
	psp.Identifier.Service = standard.GetService().Identifier.Service
	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&psp, psp.Details.Keyword, psp.Identifier.Service, psp.Identifier.User) != nil {
		_ = insert.Permission(&psp)
	}

	pss := standard.GetPermissionServiceServices()
	pss.Identifier.User = standard.GetUser().Identifier.User
	pss.Identifier.Service = standard.GetService().Identifier.Service
	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&pss, pss.Details.Keyword, pss.Identifier.Service, pss.Identifier.User) != nil {
		_ = insert.Permission(&pss)
	}

	psu := standard.GetPermissionServiceUsers()
	psu.Identifier.User = standard.GetUser().Identifier.User
	psu.Identifier.Service = standard.GetService().Identifier.Service
	if fetch.PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(&psu, psu.Details.Keyword, psu.Identifier.Service, psu.Identifier.User) != nil {
		_ = insert.Permission(&psu)
	}
}
