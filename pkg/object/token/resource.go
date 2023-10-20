package token

type Resource struct {
	Identifier string `bson:"identifier"`

	Owner bool `bson:"owner"`
	User  bool `bson:"user"`
}
