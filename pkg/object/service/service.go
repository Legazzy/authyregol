package service

type Service struct {
	Description Description `bson:"description"`
	Identifier  Identifier  `bson:"identifier"`
}
