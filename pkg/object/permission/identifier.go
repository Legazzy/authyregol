package permission

type Identifier struct {
	Permission string `bson:"permission"`
	Service    string `bson:"service"`
	User       string `bson:"user"`
}
