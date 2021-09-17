package convert

import (
	fileconfig "github.com/myeung18/service-binding-client/pkg/binding/internal/fileconfig"
	"strings"
)

// Converter converts and returns a ServiceBinding object to a database specific connection string
type Converter interface {
	Convert(sb fileconfig.ServiceBinding) string
}

type MongoDBConverter struct{}

func (m *MongoDBConverter) Convert(binding fileconfig.ServiceBinding) string {
	prefix := "mongodb://"
	if binding.Properties["srv"] == "true" {
		prefix = "mongodb+srv://"
	}

	database := ""
	if binding.Properties["options"] != "" {
		database = "?" + binding.Properties["options"]
	}
	if binding.Properties["database"] != "" {
		database = "/" + binding.Properties["database"] + database
	} else if binding.Properties["options"] != "" {
		database = "/" + database
	}

	return strings.Join([]string{prefix,
		binding.Properties["username"], ":",
		binding.Properties["password"], "@",
		binding.Properties["host"],
		database}, "")
}

// GetMongodbConnectionString returns mongoDB connection info. in a formatted string
func GetMongodbConnectionString(bindingType string) (string, error) {
	//get the binding available from file system
	bindingFileReader := fileconfig.NewBindingReader()
	serviceBindings, err := bindingFileReader.ReadServiceBindingConfig()
	if err != nil {
		return "", err
	}
	mongoBinding := singleMatchingByType(bindingType, serviceBindings)
	converter := MongoDBConverter{}
	return converter.Convert(mongoBinding), nil
}

//
func matchingByType(bindingType string, serviceBindings []fileconfig.ServiceBinding) []fileconfig.ServiceBinding {
	var res []fileconfig.ServiceBinding
	if len(serviceBindings) == 0 {
		return res
	}
	for _, sb := range serviceBindings {
		if sb.BindingType == bindingType {
			res = append(res, sb)
		}
	}
	return res
}

func singleMatchingByType(bindingType string, serviceBindings []fileconfig.ServiceBinding) fileconfig.ServiceBinding {
	res := fileconfig.ServiceBinding{}
	if len(serviceBindings) == 0 {
		return res
	}
	matchingBindings := matchingByType(bindingType, serviceBindings)
	if len(matchingBindings) == 0 {
		return res
	}
	//return the first one - should we return the 1st one after sort, or other selection criteria?
	return matchingBindings[0]
}

//PostgreSQL
