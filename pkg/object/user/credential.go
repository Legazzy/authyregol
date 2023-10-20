package user

type Credential struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Prefixes string `bson:"prefixes"`
}
