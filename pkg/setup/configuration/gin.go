package configuration

import (
	"github.com/Authyre/authyreapi/api"
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
