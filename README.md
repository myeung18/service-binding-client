# Service Binding Client for Service Binding Operator

<a href="https://github.com/myeung18/service-binding-client/actions?query=workflow%3Aunit-tests"><img alt="service-binding-client unit tests status" src="https://github.com/myeung18/service-binding-client/workflows/unit-tests/badge.svg"></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/myeung18/service-binding-client)](https://goreportcard.com/report/github.com/myeung18/service-binding-client)


Install

```shell
go get -u github.com/myeung18/service-binding-client
```

then include the client in your code
```go
import (
    github.com/myeung18/service-binding-client/pkg/binding/convert
)

// call
string, err := convert.GetMongodbConnectionString("mongodb")
if err != nil {
    fmt.Println(err)
}
fmt.Println(string)
```
  
run locally
```
SERVICE_BINDING_ROOT=bindings go run ./<main.go>
```