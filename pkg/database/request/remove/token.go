package remove

import (
	"context"
	"github.com/Authyre/authyreapi/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func TokenByIdentifierToken(token string) error {
	que := bson.M{"identifier.token": token}

	if col, err := mongo.LoadCollection(mongo.CollectionToken); err != nil {
		return err
	} else if _, err := col.DeleteOne(context.TODO(), que); err != nil {
		return err
	}

	return nil
}

func TokensByIdentifierPermission(permission string) error {
	que := bson.M{"identifier.permissions": permission}

	if col, err := mongo.LoadCollection(mongo.CollectionToken); err != nil {
		return err
	} else if _, err := col.DeleteMany(context.TODO(), que); err != nil {
		return err
	}

	return nil
}
