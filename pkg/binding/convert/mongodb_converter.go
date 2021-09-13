package convert

type ServiceBindingConfigSource struct {
	name        string
	propertyMap map[string]string
}

type Converter interface {
	Convert() ServiceBindingConfigSource
}

//MongoDB
func GetMongodbConnectionString() (string, error) {

	return "", nil
}

//PostgreSQL
