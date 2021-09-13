package read

import "os"

var (
	bindingRootEnvVar = os.Getenv("SERVICE_BINDING_ROOT")
)

func GetBindingRootDirectory() string {
	//TODO
	if bindingRootEnvVar == "" {
		return "configs/service-binding"
	}
	return bindingRootDirectory
}
