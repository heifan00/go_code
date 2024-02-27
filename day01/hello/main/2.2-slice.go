package main

import "unsafe"

func main() {

	s := *new([]int)
	m := *new(map[string]int)
	c := *new(chan int)

	println(unsafe.Sizeof(s))
	println(unsafe.Sizeof(m))
	println(unsafe.Sizeof(c))
}
