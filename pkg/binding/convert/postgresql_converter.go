package convert

import (
	"github.com/myeung18/service-binding-client/internal/fileconfig"
	"strings"
)

const (
	optionsSeparator        = "&" // separator for each option
	optionKeyValueSeparator = "=" // separator for key/value in each option
)

var escaper = strings.NewReplacer(` `, `\ `, `'`, `\'`, `\`, `\\`)

// PostgreSQLConnectionStringConverter converts a ServiceBinding object to a DSN connection string
type PostgreSQLConnectionStringConverter struct{}

// Convert converts a ServiceBinding to a DSN string
func (m *PostgreSQLConnectionStringConverter) Convert(binding fileconfig.ServiceBinding) string {
	var parts []string
	addToParts := func(k, v string) {
		parts = append(parts, k+"="+escaper.Replace(v))
	}
	addToParts("host", binding.Properties[keyHost])
	if binding.Properties[keyPort] != "" {
		addToParts("port", binding.Properties[keyPort])
	}
	if binding.Properties[keyUsername] != "" {
		addToParts("user", binding.Properties[keyUsername])
	}
	if binding.Properties[keyPassword] != "" {
		addToParts("password", binding.Properties[keyPassword])
	}
	addToParts("dbname", binding.Properties[keyDatabase])

	if binding.Properties[keyOptions] != "" {
		// multiple options are expected to be separated by & (ampersand) sign, e.g. sslmode=disable&connect_timeout=10
		options := strings.Split(binding.Properties[keyOptions], optionsSeparator)
		for _, option := range options {
			optionKeyValue := strings.Split(option, optionKeyValueSeparator)
			if len(optionKeyValue) == 2 && len(optionKeyValue[0]) > 0 && len(optionKeyValue[1]) > 0 {
				addToParts(optionKeyValue[0], optionKeyValue[1])
			}
		}
	}

	return strings.Join(parts, " ")
}
