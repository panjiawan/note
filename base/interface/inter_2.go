package main

import "fmt"

// 类型断言基本用法
func processEmptyInterface(value interface{}) {
	// 类型断言
	if str, ok := value.(string); ok {
		fmt.Printf("字符串值: %s\n", str)
	} else if num, ok := value.(int); ok {
		fmt.Printf("整数值: %d\n", num)
	} else {
		fmt.Printf("其他类型: %T\n", value)
	}

	// 类型switch
	switch v := value.(type) {
	case string:
		fmt.Printf("这是字符串: %s\n", v)
	case int:
		fmt.Printf("这是整数: %d\n", v)
	case float64:
		fmt.Printf("这是浮点数: %f\n", v)
	case bool:
		fmt.Printf("这是布尔值: %t\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

// 业务场景：处理多种类型的API响应
type APIResponse struct {
	Data    interface{}
	Code    int
	Message string
}

func handleAPIResponse(response APIResponse) {
	// 根据不同类型的数据进行不同处理
	switch data := response.Data.(type) {
	case map[string]interface{}:
		fmt.Println("处理对象数据:")
		for key, value := range data {
			fmt.Printf("  %s: %v\n", key, value)
		}
	case []interface{}:
		fmt.Printf("处理数组数据，长度: %d\n", len(data))
		for i, item := range data {
			fmt.Printf("  [%d]: %v\n", i, item)
		}
	case string:
		fmt.Printf("处理字符串数据: %s\n", data)
	case float64:
		fmt.Printf("处理数值数据: %.2f\n", data)
	default:
		fmt.Printf("未知数据类型: %T\n", data)
	}
}

func main() {
	// 测试类型处理
	processEmptyInterface("Hello")
	processEmptyInterface(42)
	processEmptyInterface(3.14)
	processEmptyInterface(true)

	// 模拟API响应处理
	responses := []APIResponse{
		{Data: map[string]interface{}{"name": "张三", "age": 25}, Code: 200},
		{Data: []interface{}{"商品1", "商品2", "商品3"}, Code: 200},
		{Data: "成功消息", Code: 200},
	}

	for _, resp := range responses {
		handleAPIResponse(resp)
		fmt.Println()
	}
}
