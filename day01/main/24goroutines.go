package main

import (
	"fmt"
	"time"
)

// 协程(goroutine) 是轻量级的执行线程。

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 假设我们有一个函数叫做 f(s)。 我们一般会这样 同步地 调用它
	f("direct")

	// 使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
	go f("goroutine")

	// 也可以为匿名函数调用启动一个 goroutine。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
