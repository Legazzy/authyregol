package remove

import (
	"context"
	"github.com/authyre/authyre-api/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func ServiceByIdentifierService(service string) error {
	que := bson.M{"identifier.service": service}

	if col, err := mongo.LoadCollection(mongo.CollectionService); err != nil {
		return err
	} else if _, err := col.DeleteOne(context.TODO(), que); err != nil {
		return err
	}

	go func(service string) { _ = PermissionsByIdentifierService(service) }(service)

	return nil
}
