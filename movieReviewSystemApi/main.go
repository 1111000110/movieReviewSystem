package main

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"log"
)

func main() {
	pusher := kq.NewPusher([]string{
		"localhost:29092",
	}, "chatSend")

	if err := pusher.Push(context.Background(), "foo"); err != nil {
		log.Fatal(err)
	}
}
