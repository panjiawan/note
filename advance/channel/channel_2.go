package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. select多路复用
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("收到:", msg1)
		case msg2 := <-ch2:
			fmt.Println("收到:", msg2)
		case <-time.After(3 * time.Second): // 超时控制
			fmt.Println("超时!")
			return
		}
	}

	// 2. 定时器与Ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("定时触发 at", t.Format("15:04:05"))
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker停止")

	// 3. 工作池模式
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for i := 0; i <= 5; i++ {
		value := <-results

		fmt.Printf("Worker 处理结果 %d\n", value)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d 处理任务 %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
