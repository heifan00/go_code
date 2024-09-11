package main

import (
	"fmt"
	"time"
)

/*通道同步*/
// 我们将要在协程中运行这个函数。 done 通道将被用于通知其他协程这个函数已经完成工作。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们已经完工啦。
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	<-done
}
