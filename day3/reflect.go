package main

import (
	"fmt"
	"reflect"
)

type myInt int

func main()  {
	var f float64 = 3.14
	var i myInt = 10
	v1 := reflect.ValueOf(f)
	v2 := reflect.ValueOf(i)
	fmt.Println("Type :", v1.Type(), "\nKind :", v1.Kind())
	fmt.Println("\nType :", v2.Type(), "\nKind :", v2.Kind())
}
