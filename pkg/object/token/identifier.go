package token

type Identifier struct {
	Token string `bson:"token"`
	User  string `bson:"user"`

	Permissions []string `bson:"permissions"`
}
