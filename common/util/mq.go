package util

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
)

func NewProducer(addr, groupName string) rocketmq.Producer {
	address, err := primitive.NewNamesrvAddr(addr)
	if err != nil {
		log.Fatal(err)
	}
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(address),
		producer.WithGroupName(groupName),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}
	if err = p.Start(); err != nil {
		panic(err)
	}
	return p
}
