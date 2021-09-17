package main

import (
	"fmt"
	"github.com/myeung18/service-binding-client/pkg/binding/convert"
)

func main() {

	conStr, err := convert.GetMongodbConnectionString("mongodb")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conStr)
}
