package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 1. 使用context控制Goroutine生命周期
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: 收到停止信号\n", id)
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Worker %d: 正在工作...\n", id)
		}
	}
}

// 2. 限制并发数量的Goroutine池
func limitedWorkerPool(tasks []string, maxConcurrency int) {
	semaphore := make(chan struct{}, maxConcurrency) // 信号量
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)
		semaphore <- struct{}{} // 获取信号量

		go func(taskNum int, taskDesc string) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			fmt.Printf("任务 %d: 开始处理 %s\n", taskNum, taskDesc)
			time.Sleep(time.Second)
			fmt.Printf("任务 %d: 完成处理 %s\n", taskNum, taskDesc)
		}(i, task)
	}

	wg.Wait()
	fmt.Println("所有任务完成!")
}

// 3. 获取Goroutine信息
func showGoroutineInfo() {
	fmt.Printf("当前Goroutine数量: %d\n", runtime.NumGoroutine())
	fmt.Printf("CPU核心数: %d\n", runtime.NumCPU())

	// 设置最大并行度
	runtime.GOMAXPROCS(4)
}

func main() {
	fmt.Println("=== Goroutine控制示例 ===")

	// 示例1: 使用context控制
	fmt.Println("\n1. Context控制:")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(4 * time.Second)

	// 示例2: 限制并发数
	fmt.Println("\n2. 限制并发数:")
	tasks := []string{
		"数据导入", "报表生成", "缓存预热",
		"数据备份", "日志归档", "用户通知",
	}
	limitedWorkerPool(tasks, 2)

	// 示例3: 系统信息
	fmt.Println("\n3. 系统信息:")
	showGoroutineInfo()
}
