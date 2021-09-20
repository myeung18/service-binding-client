package convert

import (
	"github.com/myeung18/service-binding-client/internal/fileconfig"
	"net/url"
	"strings"
)

const (
	//SpecialChars special chars used in username or password
	SpecialChars = ":/?#[]@"
	// KeyDatabase database instance name
	KeyDatabase = "database"
	// KeyHost DB host
	KeyHost = "host"
	// KeyOptions Connection options
	KeyOptions = "options"
	// KeyUsername DB User name
	KeyUsername = "username"
	//KeyPassword  DB User Password
	KeyPassword = "password"
	// KeySrv true to use DNS seed list
	KeySrv = "srv"
)

// Converter converts and returns a ServiceBinding object to a database specific connection string
type Converter interface {
	// Convert converts a ServiceBinding obj to a connection string
	Convert(sb fileconfig.ServiceBinding) string
}

// MongoDBConverter type for MongoDB
type MongoDBConverter struct{}

// Convert converts a mongodb ServiceBinding to a connection string
func (m *MongoDBConverter) Convert(binding fileconfig.ServiceBinding) string {
	prefix := "mongodb://"
	if strings.EqualFold(binding.Properties[KeySrv], "true") {
		prefix = "mongodb+srv://"
	}

	database := ""
	if binding.Properties[KeyOptions] != "" {
		database = "?" + binding.Properties[KeyOptions]
	}
	if binding.Properties[KeyDatabase] != "" {
		database = "/" + binding.Properties[KeyDatabase] + database
	} else if binding.Properties[KeyOptions] != "" {
		database = "/" + database
	}

	return strings.Join([]string{prefix,
		encodeIfContainsSpecialCharacters(binding.Properties[KeyUsername]), ":",
		encodeIfContainsSpecialCharacters(binding.Properties[KeyPassword]), "@",
		binding.Properties[KeyHost],
		database}, "")
}

func encodeIfContainsSpecialCharacters(userNameOrPassword string) string {
	if strings.ContainsAny(userNameOrPassword, SpecialChars) {
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
