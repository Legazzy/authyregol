package transfer

import "github.com/Authyre/authyreapi/pkg/object/token"

type Token struct {
	Endpoint token.Endpoint
	Resource token.Resource

	Permissions []struct {
		Keyword string
		Service string
	}
}
