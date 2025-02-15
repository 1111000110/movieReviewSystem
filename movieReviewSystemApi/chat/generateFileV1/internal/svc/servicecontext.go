package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/rest"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/hub"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/middleware"
	"movieReviewSystem/movieReviewSystemApi/chat/generateFileV1/internal/types"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	Middleware rest.Middleware
	Hub        hub.Hub
	Pusher     kq.Pusher
	Receive    queue.MessageQueue
}

func NewServiceContext(c config.Config) *ServiceContext {
	serviceContext := &ServiceContext{
		Config:     c,
		Middleware: middleware.NewMiddleware().Handle,
		Hub:        hub.NewHub(),
		Pusher: *kq.NewPusher(append([]string(nil),
			c.Pusher.Brokers...,
		), c.Pusher.Topic),
	}
	queueHandler := func(ctx context.Context, key, value string) error {
		var message types.Message
		if err := json.Unmarshal([]byte(value), &message); err != nil {
			fmt.Println("unmarshal err:", err)
		}
		fmt.Println(time.Now().Unix()) //zhangxuan 2
		serviceContext.Hub.Broadcast <- message
		return nil
	}
	serviceContext.Receive = kq.MustNewQueue(c.Receive, kq.WithHandle(queueHandler))
	return serviceContext
}
