package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// 用户注册验证函数
func validateUser(username, password string) (bool, error) {
	if len(username) < 3 {
		return false, errors.New("用户名长度不能少于3个字符")
	}
	if len(password) < 6 {
		return false, errors.New("密码长度不能少于6个字符")
	}
	return true, nil
}

// 多返回值函数：数据库查询模拟
func queryUser(userId int) (string, int, error) {
	if userId <= 0 {
		return "", 0, errors.New("无效的用户ID")
	}
	// 模拟数据库查询
	users := map[int]struct {
		name string
		age  int
	}{
		1: {"张三", 25},
		2: {"李四", 30},
	}

	if user, ok := users[userId]; ok {
		return user.name, user.age, nil
	}
	return "", 0, errors.New("用户不存在")
}

// 闭包：计数器
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// 用户注册验证
	valid, err := validateUser("ab", "123456")
	if err != nil {
		fmt.Println("验证失败:", err)
	} else {
		fmt.Println("验证通过:", valid)
	}

	// 多返回值
	name, age, err := queryUser(1)
	if err != nil {
		fmt.Println("查询失败:", err)
	} else {
		fmt.Printf("用户: %s, 年龄: %d\n", name, age)
	}

	// 闭包使用
	counter := createCounter()
	fmt.Println("计数:", counter()) // 1
	fmt.Println("计数:", counter()) // 2

	// 匿名函数：字符串处理
	strs := []string{"apple", "Banana", "cherry"}
	fmt.Println("原始:", strs)

	// 使用匿名函数排序（忽略大小写）
	sort.Slice(strs, func(i, j int) bool {
		return strings.ToLower(strs[i]) < strings.ToLower(strs[j])
	})
	fmt.Println("排序后:", strs)
}
