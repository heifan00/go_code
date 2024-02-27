package main

import (
	"fmt"
	"reflect"
)

type X = int

func main() {

	var a int = 100
	var b X = a

	fmt.Printf("%T, %v\n", b, b)
	fmt.Println(reflect.TypeOf(b))
}
