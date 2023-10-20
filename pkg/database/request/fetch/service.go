package fetch

import (
	"context"
	"github.com/authyre/authyre-api/pkg/database/mongo"
	"github.com/authyre/authyre-api/pkg/object/service"
	"go.mongodb.org/mongo-driver/bson"
)

func ServiceByIdentifierService(entity *service.Service, service string) error {
	que := bson.M{"identifier.service": service}

	if col, err := mongo.LoadCollection(mongo.CollectionService); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func ServiceByDescriptionName(entity *service.Service, name string) error {
	que := bson.M{"description.name": name}

	if col, err := mongo.LoadCollection(mongo.CollectionService); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func Services(entities *[]service.Service) error {
	que := bson.M{}

	if col, err := mongo.LoadCollection(mongo.CollectionService); err != nil {
		return err
	} else if cur, err := col.Find(context.TODO(), que); err != nil {
		return err
	} else if err := cur.All(context.Background(), entities); err != nil {
		return err
	}

	return nil
}
