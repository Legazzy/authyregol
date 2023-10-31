package fetch

import (
	"context"
	"github.com/authyre/authyreapi/pkg/database/mongo"
	"github.com/authyre/authyreapi/pkg/object/user"
	"go.mongodb.org/mongo-driver/bson"
)

func UserByCredentialUsername(entity *user.User, username string) error {
	que := bson.M{"credential.username": username}

	if col, err := mongo.LoadCollection(mongo.CollectionUser); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func UserByIdentifierUser(entity *user.User, user string) error {
	que := bson.M{"identifier.user": user}

	if col, err := mongo.LoadCollection(mongo.CollectionUser); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func Users(entities *[]user.User) error {
	que := bson.M{}

	if col, err := mongo.LoadCollection(mongo.CollectionUser); err != nil {
		return err
	} else if cur, err := col.Find(context.TODO(), que); err != nil {
		return err
	} else if err := cur.All(context.Background(), entities); err != nil {
		return err
	}

	return nil
}
