package user

type Personal struct {
	Forename string `bson:"forename"`
	Lastname string `bson:"lastname"`
}
