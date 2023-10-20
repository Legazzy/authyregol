package insert

import (
	"context"
	"github.com/authyre/authyre-api/pkg/database/mongo"
	"github.com/authyre/authyre-api/pkg/object/service"
)

func Service(entity *service.Service) error {
	if col, err := mongo.LoadCollection(mongo.CollectionService); err != nil {
		return err
	} else if _, err := col.InsertOne(context.TODO(), entity); err != nil {
		return err
	}

	return nil
}
