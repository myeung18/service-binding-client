package fileconfig

import "os"

var bindingRootEnvVar = os.Getenv("SERVICE_BINDING_ROOT")

func GetBindingRootDirectory() string {
	if bindingRootEnvVar == "" {
		return "/bindings"
	}
	return bindingRootEnvVar
}
