package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func NewProducer(addr, groupName string) (rocketmq.Producer, error) {
	address, err := primitive.NewNamesrvAddr(addr)
	if err != nil {
		return nil, err
	}
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(address),
		producer.WithGroupName(groupName),
		producer.WithRetry(2),
	)
	if err != nil {
		return nil, err
	}
	if err = p.Start(); err != nil {
		return nil, err
	}
	return p, nil
}

func SendMsg(p rocketmq.Producer, topic string, msg interface{}) error {
	_, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: topic,
		Body:  []byte(fmt.Sprint(msg)),
	})
	return err
}
