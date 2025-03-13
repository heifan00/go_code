package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		计时器代表未来的一个事件。你告诉计时器你想等待多长时间，它会提供一个届时会收到通知的通道。此计时器将等待2秒。
	*/
	timer1 := time.NewTimer(2 * time.Second)

	// 阻塞直到计时器到期。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
