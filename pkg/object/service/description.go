package service

type Description struct {
	Author  string `bson:"author"`
	Details string `bson:"details"`
	Name    string `bson:"name"`
	Version string `bson:"version"`
}
