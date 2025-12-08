package main

import "fmt"

// 可变参数函数：计算总和
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 可变参数函数：格式化打印
func formatPrint(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func main() {
	fmt.Println("总和:", sum(1, 2, 3, 4, 5))
	fmt.Println("总和:", sum(10, 20))

	formatPrint("用户: %s, 年龄: %d, 分数: %.2f\n",
		"张三", 25, 89.5)
}
