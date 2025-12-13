package main

import (
	"fmt"
	"time"
)

// 支付接口定义
type Payment interface {
	Pay(amount float64) (string, error)
	Refund(transactionID string, amount float64) (string, error)
	Query(transactionID string) (string, error)
}

// 支付宝支付实现
type Alipay struct {
	AppID      string
	AppSecret  string
	MerchantID string
}

func (a Alipay) Pay(amount float64) (string, error) {
	// 调用支付宝SDK实现支付逻辑
	return fmt.Sprintf("ALIPAY-%d", time.Now().Unix()), nil
}

func (a Alipay) Refund(transactionID string, amount float64) (string, error) {
	// 调用支付宝退款接口
	return fmt.Sprintf("ALIPAY-REFUND-%s", transactionID), nil
}

func (a Alipay) Query(transactionID string) (string, error) {
	// 查询支付宝交易状态
	return "SUCCESS", nil
}

// 微信支付实现
type WechatPay struct {
	AppID  string
	MchID  string
	APIKey string
}

func (w WechatPay) Pay(amount float64) (string, error) {
	// 调用微信支付API
	return fmt.Sprintf("WECHAT-%d", time.Now().Unix()), nil
}

func (w WechatPay) Refund(transactionID string, amount float64) (string, error) {
	// 调用微信退款API
	return fmt.Sprintf("WECHAT-REFUND-%s", transactionID), nil
}

func (w WechatPay) Query(transactionID string) (string, error) {
	// 查询微信支付状态
	return "SUCCESS", nil
}

// 支付处理函数 - 多态体现
func ProcessPayment(p Payment, amount float64) (string, error) {
	// 统一的支付处理逻辑
	fmt.Println("开始处理支付请求...")
	transactionID, err := p.Pay(amount)
	if err != nil {
		return "", err
	}
	fmt.Printf("支付成功，交易号: %s\n", transactionID)
	return transactionID, nil
}

func main() {
	// 初始化支付方式
	alipay := Alipay{
		AppID:      "2021000123456789",
		AppSecret:  "your_app_secret",
		MerchantID: "2088xxxxxx",
	}

	wechatPay := WechatPay{
		AppID:  "wx1234567890abcdef",
		MchID:  "1230000109",
		APIKey: "your_api_key",
	}

	// 使用不同的支付方式处理支付
	transactions := []struct {
		payment Payment
		amount  float64
		name    string
	}{
		{alipay, 199.99, "支付宝"},
		{wechatPay, 299.99, "微信支付"},
	}

	for _, t := range transactions {
		fmt.Printf("\n使用%s支付:\n", t.name)
		transactionID, err := ProcessPayment(t.payment, t.amount)
		if err != nil {
			fmt.Printf("支付失败: %v\n", err)
			continue
		}
		fmt.Printf("交易完成: %s\n", transactionID)
	}
}
