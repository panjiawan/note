package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID        int
	UserID    string
	Amount    float64
	Status    string
	CreatedAt time.Time
}

func orderProducer(orderChan chan<- Order, numOrders int) {
	for i := 1; i <= numOrders; i++ {
		order := Order{
			ID:        i,
			UserID:    fmt.Sprintf("user%d", rand.Intn(100)),
			Amount:    rand.Float64() * 1000,
			Status:    "pending",
			CreatedAt: time.Now(),
		}
		orderChan <- order
		fmt.Printf("生成订单: ID=%d, 用户=%s, 金额=¥%.2f\n",
			order.ID, order.UserID, order.Amount)
		time.Sleep(time.Millisecond * 100) // 模拟生成间隔
	}
	close(orderChan)
}

func orderProcessor(orderChan <-chan Order, resultChan chan<- Order) {
	for order := range orderChan {
		// 模拟订单处理逻辑
		processingTime := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(processingTime)

		// 更新订单状态
		if order.Amount > 500 {
			order.Status = "verified" // 大额订单需要验证
		} else {
			order.Status = "completed"
		}

		resultChan <- order
	}
	time.Sleep(time.Second * 2)
	close(resultChan)
}

func orderResultCollector(resultChan <-chan Order, done chan<- bool) {
	processedCount := 0
	for order := range resultChan {
		processedCount++
		fmt.Printf("处理完成: 订单ID=%d, 状态=%s, 金额=¥%.2f\n",
			order.ID, order.Status, order.Amount)
	}
	fmt.Printf("所有订单处理完成! 总计: %d 个订单\n", processedCount)
	done <- true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 创建管道
	orderChan := make(chan Order, 10)
	resultChan := make(chan Order, 10)
	done := make(chan bool)

	// 启动订单生成器
	go orderProducer(orderChan, 20)

	// 启动多个订单处理器（工人）
	for i := 1; i <= 3; i++ {
		go orderProcessor(orderChan, resultChan)
	}

	// 启动结果收集器
	go orderResultCollector(resultChan, done)

	// 等待所有处理完成
	<-done
}
