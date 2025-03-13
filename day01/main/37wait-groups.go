package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d staring\n", id)

	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	/**
	此WaitGroup用于等待此处启动的所有goroutine完成。
	注意：如果将WaitGroup显式传递给函数，则应通过指针完成。
	*/
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
