package main

import (
	"fmt"
	"unsafe"
)

// slice map channel

func main() {
	s := *new([]int)
	m := *new(map[string]int)
	c := *new(chan int)

	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(c))
}
