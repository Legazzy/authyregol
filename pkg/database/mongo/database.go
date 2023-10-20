package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Hostname = "localhost"
var Database = "authyre"
var InetPort = "27017"
var Username = "authyre"
var Password = "authyre"

func LoadDatabase() (*mongo.Client, error) {
	adr := fmt.Sprintf("mongodb://%s:%s@%s:%s", Username, Password, Hostname, InetPort)
	opt := options.Client().ApplyURI(adr)

	ctx, cnc := context.WithTimeout(context.Background(), 5*time.Second)
	cli, err := mongo.Connect(ctx, opt)

	defer cnc()
	return cli, err
}
