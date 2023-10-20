package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var CollectionPermission = "permission"
var CollectionService = "service"
var CollectionToken = "token"
var CollectionUser = "user"

func LoadCollection(collection string) (*mongo.Collection, error) {
	cli, err := LoadDatabase()
	col := cli.Database(Database).Collection(collection)

	return col, err
}
