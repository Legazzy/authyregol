package configuration

import (
	"github.com/authyre/authyre-api/pkg/database/mongo"
	"os"
)

func LoadDatabase() {
	if val, exi := os.LookupEnv("MONGO_HOSTNAME"); exi {
		mongo.Hostname = val
	}

	if val, exi := os.LookupEnv("MONGO_INETPORT"); exi {
		mongo.InetPort = val
	}

	if val, exi := os.LookupEnv("MONGO_DATABASE"); exi {
		mongo.Database = val
	}

	if val, exi := os.LookupEnv("MONGO_USERNAME"); exi {
		mongo.Username = val
	}

	if val, exi := os.LookupEnv("MONGO_PASSWORD"); exi {
		mongo.Password = val
	}
}
