package kafka

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Goboolean/fetch-server.v1/api/model"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type SubscribeListener interface {
	OnReceiveStockAggs(name string, stock *model.StockAggregate)
}

type Subscriber struct {
	consumer *kafka.Consumer
	listener SubscribeListener

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewSubscriber(c *resolver.ConfigMap, lis SubscribeListener) (*Subscriber, error) {

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	group, err := c.GetStringKey("GROUP")
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s:%s", host, port)

	config := &kafka.ConfigMap{
		"bootstrap.servers":       address,
		"auto.offset.reset":       "earliest",
		"socket.keepalive.enable": true,
		"group.id":                group,
	}

	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	instance := &Subscriber{
		consumer: consumer,
		listener: lis,
		ctx:      ctx,
		cancel:   cancel,
	}

	go instance.subscribeMessage(ctx, &instance.wg)

	return instance, nil
}

func (s *Subscriber) subscribeMessage(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:

		}

		msg, err := s.consumer.ReadMessage(time.Second)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Debug("data was not received for 1 second")
			continue
		}

		var data model.StockAggregate
		if err := proto.Unmarshal(msg.Value, &data); err != nil {
			log.WithFields(log.Fields{
				"topic": *msg.TopicPartition.Topic,
				"data":  msg.Value,
				"error": err,
			}).Error("failed to deserialize data")
			continue
		}

		topic := strings.TrimPrefix(*msg.TopicPartition.Topic, prefix)

		s.listener.OnReceiveStockAggs(topic, &data)
	}
}

func (s *Subscriber) Close() {
	s.cancel()
	s.wg.Wait()
	//s.consumer.Close()
}

func (s *Subscriber) Ping(ctx context.Context) error {
	deadline, ok := ctx.Deadline()

	if !ok {
		return fmt.Errorf("timeout setting on ctx required")
	}

	remaining := time.Until(deadline)

	_, err := s.consumer.GetMetadata(nil, true, int(remaining.Milliseconds()))
	return err
}

// only one subscription psr instance is allowed
func (s *Subscriber) Subscribe(stock string) error {
	stock = packTopic(stock)
	log.WithField("topic", stock).Debug("data sub")
	return s.consumer.Subscribe(stock, nil)
}

func (s *Subscriber) Unsubscribe(stock string) error {
	stock = packTopic(stock)
	return s.consumer.Unsubscribe()
}
