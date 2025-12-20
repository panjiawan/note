package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	//defer cancel()
	//
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err()) // 输出 "context deadline exceeded"
	//}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 调用 cancel 时，该 ctx 以及所有从它派生的子 context 都会被取消

	go func() {
		// 模拟一个会被取消的操作
		select {
		case <-ctx.Done():
			fmt.Println("Operation canceled")
		case <-time.After(5 * time.Second):
			fmt.Println("Finished operation")
		}
	}()

	cancel() // 主动发送取消信号
	time.Sleep(100 * time.Millisecond)

	//req, err := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 构造一个超时间为50毫秒的Context
	//ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	//defer cancel()
	//req = req.WithContext(ctx)
	//
	//c := &http.Client{}
	//res, err := c.Do(req)
	//if err != nil {
	//	log.Fatal("1-", err)
	//}
	//defer res.Body.Close()
	//out, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal("2-", err)
	//}
	//log.Println(string(out))
}
