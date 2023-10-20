package insert

import (
	"context"
	"github.com/authyre/authyre-api/pkg/database/mongo"
	"github.com/authyre/authyre-api/pkg/object/permission"
)

func Permission(entity *permission.Permission) error {
	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if _, err := col.InsertOne(context.TODO(), entity); err != nil {
		return err
	}

	return nil
}