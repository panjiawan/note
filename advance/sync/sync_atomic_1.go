package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1) // 比mutex性能更好
}

func (c *Counter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	// 模拟100个用户同时点赞
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// 模拟10个用户取消点赞
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}

	wg.Wait()

	fmt.Printf("最终点赞数: %d\n", counter.Value())
}
