package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func NewPushConsumer(addr, groupName string) (rocketmq.PushConsumer, error) {
	address, err := primitive.NewNamesrvAddr(addr)
	if err != nil {
		return nil, err
	}

	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(address),
		consumer.WithGroupName(groupName),
		)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Subscribe(c rocketmq.PushConsumer, topic string, handler func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	err := c.Subscribe(topic, consumer.MessageSelector{}, handler)
	if err != nil {
		return err
	}
	if err = c.Start(); err != nil {
		return err
	}
	return nil
}
