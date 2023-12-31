package main

import (
	"context"
	"github.com/Authyre/authyreapi/api"
	"github.com/Authyre/authyreapi/pkg/database/mongo"
	"github.com/Authyre/authyreapi/pkg/setup/configuration"
	"github.com/Authyre/authyreapi/pkg/setup/population"
	"sync"
)

func main() {
	configuration.LoadDatabase()
	configuration.LoadGin()
	configuration.LoadPopulation()

	if cli, err := mongo.LoadDatabase(); err != nil {
		panic(err)
	} else {
		if err := cli.Ping(context.TODO(), nil); err != nil {
			panic(err)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() { defer wg.Done(); population.InsertServices() }()
	go func() { defer wg.Done(); population.InsertUsers() }()
	go func() { defer wg.Done(); population.InsertPermissions() }()

	wg.Wait()

	api.Attach()
}
