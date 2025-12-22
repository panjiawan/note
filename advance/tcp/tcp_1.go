package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type ChatRoom struct {
	clients map[net.Conn]string
	mutex   sync.RWMutex
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		clients: make(map[net.Conn]string),
	}
}

func (cr *ChatRoom) broadcast(sender net.Conn, message string) {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()

	for client := range cr.clients {
		if client != sender {
			client.Write([]byte(message))
		}
	}
}

func (cr *ChatRoom) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 获取用户名
	conn.Write([]byte("请输入你的用户名: "))
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("n:", n)
	//fmt.Printf("Received: %s\n", buffer[:n])
	//username := string(buffer[:n-1]) // 去掉换行符
	username := strings.TrimSpace(string(buffer[:n]))
	fmt.Printf("用户 %s 加入聊天室\n", username)

	// 注册用户
	cr.mutex.Lock()
	cr.clients[conn] = username
	cr.mutex.Unlock()

	// 广播用户加入
	cr.broadcast(conn, fmt.Sprintf("系统: %s 加入了聊天室\n", username))

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}

		message := string(buffer[:n])
		if message == "/quit\n" {
			break
		}

		// 广播消息
		fullMessage := fmt.Sprintf("%s: %s", username, message)
		cr.broadcast(conn, fullMessage)
	}

	// 用户退出
	cr.mutex.Lock()
	delete(cr.clients, conn)
	cr.mutex.Unlock()
	cr.broadcast(conn, fmt.Sprintf("系统: %s 离开了聊天室\n", username))
}

func main() {
	chatRoom := NewChatRoom()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("启动服务器失败:", err)
		return
	}
	defer listener.Close()

	fmt.Println("聊天室服务器启动在 :8888 端口...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败:", err)
			continue
		}

		go chatRoom.handleConnection(conn)
	}
}
