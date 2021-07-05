package main

import (
	"context"
	"github.com/touchtechnologies-product/message-broker"
	"github.com/touchtechnologies-product/message-broker/common"
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
	go broker.Start(func(ctx context.Context, err error) {})
	msg := []byte("Hi App2 from App1")
	err = broker.SendTopicMessage("my-topic", msg)


}