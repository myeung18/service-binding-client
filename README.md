# Service Binding Client for Service Binding Operator

Install

```shell
go get -u github.com/myeung18/service-binding-client
```

then include the client in your code
```go
import (
    github.com/myeung18/service-binding-client/pkg/binding/convert
)

# call
string, err := convert.GetMongodbConnectionString("mongodb")
if err != nil {
    fmt.Println(err)
}
fmt.Println(string)
  
# run 
SERVICE_BINDING_ROOT=bindings go run ./<main.go>
```