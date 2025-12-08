package main

import (
	"fmt"
	"strings"
)

// 定义函数类型
type StringProcessor func(string) string

// 接收函数作为参数
func processString(s string, processor StringProcessor) string {
	return processor(s)
}

func main() {
	// 函数作为参数传递
	toUpper := func(s string) string {
		return strings.ToUpper(s)
	}

	result := processString("hello go", toUpper)
	fmt.Println(result) // HELLO GO

	// 直接传递匿名函数
	result2 := processString("  hello  ", func(s string) string {
		return strings.TrimSpace(s)
	})
	fmt.Printf("'%s'\n", result2) // 'hello'
}
