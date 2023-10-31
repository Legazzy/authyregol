package configuration

import (
	"github.com/authyre/authyreapi/pkg/setup/population"
	"os"
	"strconv"
)

func LoadPopulation() {
	if val, err := strconv.ParseBool(os.Getenv("POPULATION_REPOPULATE_PERMISSIONS")); err == nil {
		population.RepopulatePermissions = val
	}

	if val, err := strconv.ParseBool(os.Getenv("POPULATION_REPOPULATE_SERVICES")); err == nil {
		population.RepopulateServices = val
	}

	if val, err := strconv.ParseBool(os.Getenv("POPULATION_REPOPULATE_USERS")); err == nil {
		population.RepopulateUsers = val
	}
}
