package standard

import "github.com/Authyre/authyreapi/pkg/object/permission"

func GetPermissionPersonalChanges() permission.Permission {
	per := permission.NewPermission()

	per.Details.Keyword = "personal_changes"
	per.Details.Description = "For applying changes to the personal information"

	return per
}

func GetPermissionPersonalInfos() permission.Permission {
	per := permission.NewPermission()

	per.Details.Keyword = "personal_infos"
	per.Details.Description = "For reading the personal information"

	return per
}

func GetPermissionPersonalPermissions() permission.Permission {
	per := permission.NewPermission()

	per.Details.Keyword = "personal_permissions"
	per.Details.Description = "For viewing and modifying the personal permissions"

	return per
}

func GetPermissionPersonalTokens() permission.Permission {
	per := permission.NewPermission()

	per.Details.Keyword = "personal_tokens"
	per.Details.Description = "For viewing and modifying the personal tokens"

	return per
}

func GetPermissionServicePermissions() permission.Permission {
	per := permission.NewPermission()

	per.Details.Keyword = "service_permissions"
	per.Details.Description = "For managing other users permissions"

	return per
}

func GetPermissionServiceServices() permission.Permission {
	per := permission.NewPermission()

	per.Details.Keyword = "service_services"
	per.Details.Description = "For managing the services of the service"

	return per
}

func GetPermissionServiceUsers() permission.Permission {
	per := permission.NewPermission()

	per.Details.Keyword = "service_users"
	per.Details.Description = "For managing the users of the service"

	return per
}
