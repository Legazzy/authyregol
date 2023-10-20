package insert

import (
	"context"
	"github.com/authyre/authyre-api/pkg/database/mongo"
	"github.com/authyre/authyre-api/pkg/object/token"
)

func Token(entity *token.Token) error {
	if col, err := mongo.LoadCollection(mongo.CollectionToken); err != nil {
		return err
	} else if _, err := col.InsertOne(context.TODO(), entity); err != nil {
		return err
	}

	return nil
}