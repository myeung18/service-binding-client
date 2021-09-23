package main

import (
	"fmt"
	"github.com/myeung18/service-binding-client/pkg/binding/convert"
)

func main() {
	connString, err := convert.GetMongoDBConnectionString()
	if err != nil {
		panic(err)
	}
	fmt.Println(connString)
	connString, err = convert.GetPostgreSQLConnectionString()
	if err != nil {
		panic(err)
	}
	fmt.Println(connString)
}
