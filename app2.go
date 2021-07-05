package main

import (
	"context"
	"fmt"
	"github.com/touchtechnologies-product/message-broker"
	"github.com/touchtechnologies-product/message-broker/common"
	"time"
)
func main(){
	conf := &common.Config{
		BackOffTime:  2,
		MaximumRetry: 3,
		Version:      "2.6.1",
		Group:        "my-group",
		Host:         []string{"localhost:9094"},
		Debug:        true,
	}
	broker, err := message.NewBroker(common.KafkaBrokerType,conf)
	if err != nil{
		panic(err)
	}
	handler := func(ctx context.Context, msg []byte) {
		fmt.Println(string(msg))
	}

	broker.RegisterHandler("my-topic", handler)
	go broker.Start(func(ctx context.Context, err error) {})
	time.Sleep(10 * time.Second)
}