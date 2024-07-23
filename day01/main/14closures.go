package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println("*******")

	test := initTest()
	fmt.Println(test())
	fmt.Println(test())
}

func initTest() func() int {
	t := 1
	return func() int {
		t += 10
		return t
	}
}
