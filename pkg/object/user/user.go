package user

type User struct {
	Address    Address    `bson:"address"`
	Contact    Contact    `bson:"contact"`
	Credential Credential `bson:"credential"`
	Identifier Identifier `bson:"identifier"`
	Personal   Personal   `bson:"personal"`
}
