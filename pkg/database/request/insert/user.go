package insert

import (
	"context"
	"github.com/Authyre/authyreapi/pkg/database/mongo"
	"github.com/Authyre/authyreapi/pkg/object/user"
)

func User(entity *user.User) error {
	if col, err := mongo.LoadCollection(mongo.CollectionUser); err != nil {
		return err
	} else if _, err := col.InsertOne(context.TODO(), entity); err != nil {
		return err
	}

	return nil
}
