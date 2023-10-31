package insert

import (
	"context"
	"github.com/Authyre/authyreapi/pkg/database/mongo"
	"github.com/Authyre/authyreapi/pkg/object/token"
)

func Token(entity *token.Token) error {
	if col, err := mongo.LoadCollection(mongo.CollectionToken); err != nil {
		return err
	} else if _, err := col.InsertOne(context.TODO(), entity); err != nil {
		return err
	}

	return nil
}
