package token

type Credential struct {
	Access  string `bson:"access"`
	Refresh string `bson:"refresh"`
}
