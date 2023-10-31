package fetch

import (
	"context"
	"github.com/authyre/authyreapi/pkg/database/mongo"
	"github.com/authyre/authyreapi/pkg/object/permission"
	"go.mongodb.org/mongo-driver/bson"
)

func PermissionByDetailKeywordAndIdentifierPermissions(entity *permission.Permission, keyword string, permissions []string) error {
	que := bson.M{"details.keyword": keyword, "identifier.permission": bson.M{"$in": permissions}}

	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func PermissionByDetailsKeywordAndIdentifierServiceAndIdentifierUser(entity *permission.Permission, keyword string, service string, user string) error {
	que := bson.M{"details.keyword": keyword, "identifier.service": service, "identifier.user": user}

	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func PermissionByIdentifierPermission(entity *permission.Permission, permission string) error {
	que := bson.M{"identifier.permission": permission}

	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if err := col.FindOne(context.TODO(), que).Decode(entity); err != nil {
		return err
	}

	return nil
}

func PermissionsByIdentifierServiceAndIdentifierUser(entities *[]permission.Permission, service string, user string) error {
	que := bson.M{"identifier.service": service, "identifier.user": user}

	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if cur, err := col.Find(context.TODO(), que); err != nil {
		return err
	} else if err := cur.All(context.TODO(), entities); err != nil {
		return err
	}

	return nil
}
