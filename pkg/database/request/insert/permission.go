package insert

import (
	"context"
	"github.com/Authyre/authyreapi/pkg/database/mongo"
	"github.com/Authyre/authyreapi/pkg/object/permission"
)

func Permission(entity *permission.Permission) error {
	if col, err := mongo.LoadCollection(mongo.CollectionPermission); err != nil {
		return err
	} else if _, err := col.InsertOne(context.TODO(), entity); err != nil {
		return err
	}

	return nil
}
