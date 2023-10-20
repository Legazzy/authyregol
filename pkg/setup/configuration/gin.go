package configuration

import (
	"github.com/authyre/authyre-api/api"
	"os"
)

func LoadGin() {
	if val, exi := os.LookupEnv("API_HOSTNAME"); exi {
		api.Hostname = val
	}

	if val, exi := os.LookupEnv("API_PORT"); exi {
		api.InetPort = val
	}
}
