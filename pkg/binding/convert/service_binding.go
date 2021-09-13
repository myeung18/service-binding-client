package convert

type ServiceBinding struct {
	Name        string
	Provider    string
	Properties  map[string]string
	BindingType string
}

func NewServiceBinding() *ServiceBinding {
	return nil
}

func (sb *ServiceBinding) MatchingByType(bindingType string) []ServiceBinding {
	binding := []ServiceBinding{{}}
	return binding
}

func (sb *ServiceBinding) singleMatchingByType(bindingType string) *ServiceBinding {
	binding := &ServiceBinding{}
	return binding
}
