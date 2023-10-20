package fetch

import (
	"context"
	"github.com/authyre/authyre-api/pkg/database/mongo"
	"github.com/authyre/authyre-api/pkg/object/token"
	"go.mongodb.org/mongo-driver/bson"
)

func TokenByCredentialAccess(entity *token.Token, access string) error {
	que := bson.M{"credential.access": access}

	if col, err := mongo.LoadCollection(mongo.CollectionToken); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func TokensByIdentifierUser(entities *[]token.Token, user string) error {
	que := bson.M{"identifier.user": user}

	if col, err := mongo.LoadCollection(mongo.CollectionToken); err != nil {
		return err
	} else if cur, err := col.Find(context.TODO(), que); err != nil {
		return err
	} else if err := cur.All(context.Background(), entities); err != nil {
		return err
	}

	return nil
}
