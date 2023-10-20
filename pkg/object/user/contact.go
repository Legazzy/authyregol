package user

type Contact struct {
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
	Website   string `bson:"website"`
	Instagram string `bson:"instagram"`
	Facebook  string `bson:"facebook"`
	Twitter   string `bson:"twitter"`
}
