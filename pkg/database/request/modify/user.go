package modify

import (
	"context"
	"github.com/authyre/authyreapi/pkg/database/mongo"
	"github.com/authyre/authyreapi/pkg/object/user"
	"go.mongodb.org/mongo-driver/bson"
)

func UserByIdentifierUser(entity *user.User, user string) error {
	que := bson.M{"identifier.user": user}

	if col, err := mongo.LoadCollection(mongo.CollectionUser); err != nil {
		return err
	} else if _, err := col.ReplaceOne(context.TODO(), que, entity); err != nil {
		return err
	}

	return nil
}
