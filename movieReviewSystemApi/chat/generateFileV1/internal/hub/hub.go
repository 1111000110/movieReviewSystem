package hub

import (
	"fmt"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/types"
)

type Hub struct {
	Clients    map[int64]Client   // 存储已注册的客户端,键为用户id，值为对应客户端
	Register   chan Client        // 注册请求的通道，用于接收客户端的注册请求
	Unregister chan Client        // 注销请求的通道，用于接收客户端的注销请求
	Broadcast  chan types.Message // 广播消息的通道，用于接收客户端发送的消息
}

// NewHub 创建一个新的 Hub 实例。
func NewHub() Hub {
	return Hub{
		Register:   make(chan Client),               // 初始化注册请求通道
		Unregister: make(chan Client),               // 初始化注销请求通道
		Clients:    make(map[int64]Client),          // 初始化客户端集合
		Broadcast:  make(chan types.Message, 10000), // 带缓冲，消息堆积成啥了，草
	}
}

// Run 启动 Hub 的主循环，处理注册、注销和消息发送。
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register: // 处理客户端注册请求
			h.Clients[client.GetClientId()] = client // 将客户端添加到客户端集合中
		case client := <-h.Unregister: // 处理客户端注销请求
			if _, ok := h.Clients[client.GetClientId()]; ok {
				delete(h.Clients, client.GetClientId()) // 如果客户端已注册，则从集合中删除
				close(client.GetSendBuffer())           // 关闭客户端的发送通道
			}
		case message := <-h.Broadcast: //有消息
			fmt.Println(message)
			if client, ok := h.Clients[message.ToUserId]; ok {
				select {
				case client.GetSendBuffer() <- message: // 将消息发送到客户端的发送通道
				default: // 如果发送失败（通道已满），关闭客户端的发送通道并从集合中删除
					close(client.GetSendBuffer())
					delete(h.Clients, client.GetClientId())
				}
			}
		}
		fmt.Println(len(h.Broadcast), len(h.Register), len(h.Unregister), len(h.Unregister))
	}
}
