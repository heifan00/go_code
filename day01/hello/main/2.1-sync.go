package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var x atomic.Int64

	go func() {
		x.Add(1)
	}()

	x.Store(x.Load() + 1)

	fmt.Println(x)

}
