package remove

import (
	"context"
	"github.com/authyre/authyreapi/pkg/database/mongo"
	"github.com/authyre/authyreapi/pkg/object/permission"
	"go.mongodb.org/mongo-driver/bson"
)

func PermissionByIdentifierPermission(permission string) error {
	que := bson.M{"identifier.permission": permission}

	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if _, err := col.DeleteOne(context.TODO(), que); err != nil {
		return err
	}

	if err := TokensByIdentifierPermission(permission); err != nil {
		return err
	}

	return nil
}

func PermissionsByIdentifierService(service string) error {
	que := bson.M{"identifier.service": service}

	var tar []permission.Permission
	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if cur, err := col.Find(context.TODO(), que); err != nil {
		return err
	} else if err := cur.All(context.TODO(), &tar); err != nil {
		return err
	} else if _, err := col.DeleteMany(context.TODO(), que); err != nil {
		return err
	}

	for _, one := range tar {
		go func(permission string) { _ = TokensByIdentifierPermission(permission) }(one.Identifier.Permission)
	}

	return nil
}

func PermissionsByIdentifierUser(user string) error {
	que := bson.M{"identifier.user": user}

	var tar []permission.Permission
	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if cur, err := col.Find(context.TODO(), que); err != nil {
		return err
	} else if err := cur.All(context.TODO(), &tar); err != nil {
		return err
	} else if _, err := col.DeleteMany(context.TODO(), que); err != nil {
		return err
	}

	for _, one := range tar {
		go func(permission string) { _ = TokensByIdentifierPermission(permission) }(one.Identifier.Permission)
	}

	return nil
}
