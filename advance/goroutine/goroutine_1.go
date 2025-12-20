package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 模拟网络请求
func fetchData(url string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done() // 完成时通知WaitGroup
	}

	// 模拟网络延迟
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)

	fmt.Printf("从 %s 获取数据完成，耗时 %v\n", url, delay)
}

func main() {
	fmt.Println("=== 串行请求 ===")
	start := time.Now()

	// 串行执行
	fetchData("https://api.example.com/users", nil)
	fetchData("https://api.example.com/products", nil)
	fetchData("https://api.example.com/orders", nil)

	fmt.Printf("串行总耗时: %v\n\n", time.Since(start))

	fmt.Println("=== 并发请求 ===")
	start = time.Now()

	// 使用WaitGroup等待所有Goroutine完成
	var wg sync.WaitGroup
	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
		"https://api.example.com/comments",
		"https://api.example.com/reviews",
	}

	// 并发执行
	for _, url := range urls {
		wg.Add(1) // 计数器+1
		go fetchData(url, &wg)
	}

	wg.Wait() // 等待所有Goroutine完成
	fmt.Printf("并发总耗时: %v\n", time.Since(start))
}
