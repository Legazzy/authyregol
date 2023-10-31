package token

import (
	"github.com/authyre/authyreapi/pkg/setup/tool"
	"github.com/google/uuid"
	"time"
)

func NewToken() Token {
	tok := Token{}

	tok.Credential.Access = tool.NewHash(uuid.New().String())
	tok.Credential.Refresh = tool.NewHash(uuid.New().String())
	tok.Identifier.Token = uuid.New().String()
	tok.Resource.User = true
	tok.Resource.Owner = true
	tok.Timestamp.Creation = time.Now().Unix()
	tok.Timestamp.Expiration = time.Now().Add(time.Hour * 24 * 30).Unix()

	return tok
}
