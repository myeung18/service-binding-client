package convert

import (
	fileconfig "github.com/myeung18/service-binding-client/pkg/binding/internal/fileconfig"
	"net/url"
	"strings"
)

const (
	SPECIAL_CHARS = ":/?#[]@"
	KEY_DATABASE  = "database"
	KEY_HOST      = "host"
	KEY_OPTIONS   = "options"
	KEY_USERNAME  = "username"
	KEY_PASSWORD  = "password"
	KEY_SRV       = "srv"
)

// Converter converts and returns a ServiceBinding object to a database specific connection string
type Converter interface {
	Convert(sb fileconfig.ServiceBinding) string
}

type MongoDBConverter struct{}

func (m *MongoDBConverter) Convert(binding fileconfig.ServiceBinding) string {
	prefix := "mongodb://"
	if strings.EqualFold(binding.Properties[KEY_SRV], "true") {
		prefix = "mongodb+srv://"
	}

	database := ""
	if binding.Properties[KEY_OPTIONS] != "" {
		database = "?" + binding.Properties[KEY_OPTIONS]
	}
	if binding.Properties[KEY_DATABASE] != "" {
		database = "/" + binding.Properties[KEY_DATABASE] + database
	} else if binding.Properties[KEY_OPTIONS] != "" {
		database = "/" + database
	}

	return strings.Join([]string{prefix,
		encodeIfContainsSpecialCharacters(binding.Properties[KEY_USERNAME]), ":",
		encodeIfContainsSpecialCharacters(binding.Properties[KEY_PASSWORD]), "@",
		binding.Properties[KEY_HOST],
		database}, "")
}

func encodeIfContainsSpecialCharacters(userNameOrPassword string) string {
	if strings.ContainsAny(userNameOrPassword, SPECIAL_CHARS) {
		return url.QueryEscape(userNameOrPassword)
	}
	return userNameOrPassword
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
