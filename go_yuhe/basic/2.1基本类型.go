package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	a, b, c := 0b1010, 0o144, 0x64

	fmt.Printf("0b%b, %#o, %#x\n", a, b, c)
	fmt.Println(math.MinInt8, math.MaxInt8)

	d := 11_22_33_44_55_66
	println(d)

	//var m map[string]int = nil
	//var c []int = nil
	//fmt.Println(m == c)

	type X = int // 别名

	var z int = 100
	var q X = z

	fmt.Printf("%T, %v\n", q, q)
	fmt.Println(reflect.TypeOf(q))

}
