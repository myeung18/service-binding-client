package fileconfig

type ServiceBinding struct {
	Name        string
	Provider    string
	Properties  map[string]string
	BindingType string
}

