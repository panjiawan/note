package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 订单结构
type Order struct {
	ID       int
	UserID   int
	Amount   float64
	Status   string
	CreateAt time.Time
}

// 1. 生产者-消费者：订单生成与处理
func orderProducer(orderCh chan<- Order, count int) {
	defer close(orderCh)
	for i := 1; i <= count; i++ {
		order := Order{
			ID:       i,
			UserID:   rand.Intn(1000) + 1,
			Amount:   rand.Float64() * 1000,
			Status:   "pending",
			CreateAt: time.Now(),
		}
		fmt.Printf("生成订单: ID=%d, 金额=¥%.2f\n", order.ID, order.Amount)
		orderCh <- order
		time.Sleep(100 * time.Millisecond) // 模拟生成间隔
	}
}

func orderConsumer(orderCh <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderCh {
		// 模拟订单处理
		time.Sleep(200 * time.Millisecond)
		order.Status = "processed"
		fmt.Printf("处理订单: ID=%d, 状态=%s\n", order.ID, order.Status)
	}
}

// 2. 扇出模式：一个订单流分发给多个处理器
func fanOutProcessor(input <-chan Order, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range input {
		// 不同的处理器做不同的处理
		switch workerID {
		case 1:
			// 处理器1：计算折扣
			discount := order.Amount * 0.1
			fmt.Printf("处理器%d计算折扣: 订单%d 优惠¥%.2f\n",
				workerID, order.ID, discount)
		case 2:
			// 处理器2：发送通知
			fmt.Printf("处理器%d发送通知: 订单%d创建成功\n",
				workerID, order.ID)
		case 3:
			// 处理器3：记录日志
			fmt.Printf("处理器%d记录日志: 订单%d金额¥%.2f\n",
				workerID, order.ID, order.Amount)
		}
		time.Sleep(150 * time.Millisecond)
	}
}

// 3. 扇入模式：多个数据源合并
func fanInProducer(producerID int, output chan<- Order) {
	defer fmt.Printf("生产者%d结束\n", producerID)
	for i := 1; i <= 3; i++ {
		order := Order{
			ID:     producerID*100 + i,
			UserID: producerID,
			Amount: float64(producerID*100 + i),
			Status: "new",
		}
		fmt.Printf("生产者%d生成订单: %d\n", producerID, order.ID)
		output <- order
		time.Sleep(time.Duration(producerID) * 100 * time.Millisecond)
	}
}

// 4. Pipeline模式：订单处理流水线
func validationStage(input <-chan Order) <-chan Order {
	output := make(chan Order, 10)
	go func() {
		defer close(output)
		for order := range input {
			// 第一阶段：订单验证
			time.Sleep(50 * time.Millisecond)
			if order.Amount > 0 {
				order.Status = "validated"
				fmt.Printf("验证通过: 订单%d\n", order.ID)
				output <- order
			} else {
				fmt.Printf("验证失败: 订单%d金额异常\n", order.ID)
			}
		}
	}()
	return output
}

func paymentStage(input <-chan Order) <-chan Order {
	output := make(chan Order, 10)
	go func() {
		defer close(output)
		for order := range input {
			// 第二阶段：支付处理
			time.Sleep(100 * time.Millisecond)
			order.Status = "paid"
			fmt.Printf("支付成功: 订单%d\n", order.ID)
			output <- order
		}
	}()
	return output
}

func shippingStage(input <-chan Order) <-chan Order {
	output := make(chan Order, 10)
	go func() {
		defer close(output)
		for order := range input {
			// 第三阶段：发货处理
			time.Sleep(150 * time.Millisecond)
			order.Status = "shipped"
			fmt.Printf("发货完成: 订单%d\n", order.ID)
			output <- order
		}
	}()
	return output
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== 1. 生产者-消费者模式演示 ===")
	// 创建订单channel
	orderCh := make(chan Order, 5)
	var wg sync.WaitGroup

	// 启动消费者
	wg.Add(2)
	go orderConsumer(orderCh, &wg)
	go orderConsumer(orderCh, &wg)

	// 启动生产者
	go orderProducer(orderCh, 6)

	wg.Wait()

	fmt.Println("\n=== 2. 扇出模式演示 ===")
	// 扇出：一个输入，多个处理器
	fanOutCh := make(chan Order, 10)
	var fanOutWg sync.WaitGroup

	// 启动3个处理器
	fanOutWg.Add(3)
	for i := 1; i <= 3; i++ {
		go fanOutProcessor(fanOutCh, i, &fanOutWg)
	}

	// 生产一些测试数据
	go func() {
		for i := 1; i <= 6; i++ {
			fanOutCh <- Order{ID: i, Amount: float64(i * 100)}
		}
		close(fanOutCh)
	}()

	fanOutWg.Wait()

	fmt.Println("\n=== 3. 扇入模式演示 ===")
	// 扇入：多个生产者，一个输出
	fanInCh := make(chan Order, 10)

	// 启动3个生产者
	for i := 1; i <= 3; i++ {
		go fanInProducer(i, fanInCh)
	}

	// 收集结果
	go func() {
		time.Sleep(2 * time.Second)
		close(fanInCh)
	}()

	fmt.Println("收集到的订单:")
	for order := range fanInCh {
		fmt.Printf("  订单ID: %d, 金额: ¥%.2f\n", order.ID, order.Amount)
	}

	fmt.Println("\n=== 4. Pipeline模式演示 ===")
	// 创建初始输入
	pipelineInput := make(chan Order, 10)

	// 构建流水线
	validatedOrders := validationStage(pipelineInput)
	paidOrders := paymentStage(validatedOrders)
	shippedOrders := shippingStage(paidOrders)

	// 发送测试订单到流水线
	go func() {
		for i := 1; i <= 3; i++ {
			pipelineInput <- Order{
				ID:     i,
				UserID: i * 10,
				Amount: float64(i * 50),
				Status: "new",
			}
		}
		close(pipelineInput)
	}()

	// 收集最终结果
	fmt.Println("流水线处理结果:")
	for order := range shippedOrders {
		fmt.Printf("  完成: 订单%d, 状态: %s\n", order.ID, order.Status)
	}

	fmt.Println("\n🎉 所有并发模式演示完成!")
}
