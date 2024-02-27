package main

import (
	"fmt"
)

func main() {
	a := 11_22_33_44
	fmt.Println(a)
	var m map[string]int = nil
	var c []int = nil
	fmt.Println(m == c)
}
