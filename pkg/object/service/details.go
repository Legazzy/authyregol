package service

type Details struct {
	Address     string `bson:"address"`
	Author      string `bson:"author"`
	Description string `bson:"description"`
	Name        string `bson:"name"`
	Version     string `bson:"version"`
}
