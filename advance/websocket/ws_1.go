package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 创建WebSocket升级器
var upgrader = websocket.Upgrader{
	// 允许所有跨域请求（生产环境应该严格限制）
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// 设置WebSocket路由
	http.HandleFunc("/ws", handleWebSocket)

	// 启动静态文件服务（用于提供HTML页面）
	http.Handle("/", http.FileServer(http.Dir("./")))

	fmt.Println("WebSocket服务器启动在 :8080")
	fmt.Println("访问 http://localhost:8080 测试聊天功能")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 1. 升级HTTP连接到WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级WebSocket失败:", err)
		return
	}
	defer conn.Close() // 确保连接最终会关闭

	fmt.Println("新的WebSocket连接建立!")

	// 2. 持续监听和处理消息
	for {
		// 读取客户端发送的消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息失败:", err)
			break
		}

		fmt.Printf("收到消息: %s\n", message)

		// 3. 向客户端回送消息
		response := fmt.Sprintf("服务器回复: 收到你的消息 '%s'", message)
		err = conn.WriteMessage(messageType, []byte(response))
		if err != nil {
			log.Println("发送消息失败:", err)
			break
		}
	}

	fmt.Println("WebSocket连接关闭")
}
