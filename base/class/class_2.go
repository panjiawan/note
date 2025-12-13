package main

import "fmt"

// 通知接口
type Notifier interface {
	Notify(message string) error
}

// 邮件通知
type EmailNotifier struct {
	SMTPHost string
	Port     int
}

func (en EmailNotifier) Notify(message string) error {
	fmt.Printf("发送邮件通知: %s\n", message)
	return nil
}

// 短信通知
type SMSNotifier struct {
	APIKey string
}

func (sn SMSNotifier) Notify(message string) error {
	fmt.Printf("发送短信通知: %s\n", message)
	return nil
}

// 应用服务
type OrderService struct {
	notifier Notifier // 组合接口
}

func (os *OrderService) SetNotifier(notifier Notifier) {
	os.notifier = notifier
}

func (os *OrderService) CreateOrder(product string, quantity int) {
	// 创建订单逻辑
	fmt.Printf("创建订单: %s x %d\n", product, quantity)

	// 发送通知
	message := fmt.Sprintf("订单创建成功: %s x %d", product, quantity)
	os.notifier.Notify(message)
}

// 广播通知 - 组合多个通知器
type BroadcastNotifier struct {
	notifiers []Notifier
}

func (bn BroadcastNotifier) Notify(message string) error {
	for _, notifier := range bn.notifiers {
		notifier.Notify(message)
	}
	return nil
}

func main() {
	orderService := &OrderService{}

	// 使用邮件通知
	emailNotifier := EmailNotifier{SMTPHost: "smtp.example.com", Port: 587}
	orderService.SetNotifier(emailNotifier)
	orderService.CreateOrder("iPhone 14", 2)

	// 切换为短信通知
	smsNotifier := SMSNotifier{APIKey: "your_api_key"}
	orderService.SetNotifier(smsNotifier)
	orderService.CreateOrder("MacBook Pro", 1)

	// 多种通知方式组合
	broadcastNotifier := BroadcastNotifier{
		notifiers: []Notifier{emailNotifier, smsNotifier},
	}
	orderService.SetNotifier(broadcastNotifier)
	orderService.CreateOrder("iPad Air", 3)
}
