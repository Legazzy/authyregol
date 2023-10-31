package remove

import (
	"context"
	"github.com/Authyre/authyreapi/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func UserByIdentifierUser(user string) error {
	que := bson.M{"identifier.user": user}

	if col, err := mongo.LoadCollection(mongo.CollectionUser); err != nil {
		return err
	} else if _, err := col.DeleteOne(context.TODO(), que); err != nil {
		return err
	}

	go func(user string) { _ = PermissionsByIdentifierUser(user) }(user)

	return nil
}
