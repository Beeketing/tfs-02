package producer

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// SimpleProducer a simple producer structure
type SimpleProducer struct {
	ctx        context.Context
	wg         *sync.WaitGroup
	channel    *amqp.Channel
	exchange   string
	exchType   string
	routingKey string
}

// NewSimpleProducer creates new producer
func NewSimpleProducer(ctx context.Context, wg *sync.WaitGroup, chann *amqp.Channel, exchange, exchType, routingKey string) *SimpleProducer {
	return &SimpleProducer{
		ctx:        ctx,
		wg:         wg,
		channel:    chann,
		exchange:   exchange,
		exchType:   exchType,
		routingKey: routingKey,
	}
}

// Start start generating data
func (p *SimpleProducer) Start() {
	if p.channel == nil || p.exchType == "" || p.exchange == "" {
		fmt.Println("Wrong producer config")
		return
	}
	// declare exchanges
	p.declare()

	// create a ticker
	ticker := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-ticker.C:
			msg := randStringRunes(100)
			fmt.Println("Sending message: ", msg)
			err := p.publish(p.exchange, p.routingKey, msg)
			if err != nil {
				fmt.Println("error when publishing data: ", err)
			}
		case <-p.ctx.Done():
			fmt.Println("Exiting consumer")
			ticker.Stop()
			p.wg.Done()
			return
		}
	}
}

func (p *SimpleProducer) publish(exch, routingKey, body string) error {
	if err := p.channel.Publish(
		exch,       // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
		},
	); err != nil {
		return fmt.Errorf("publish data error: %s", err)
	}
	return nil
}

// declare exchange and queue, also bind queue to exchange
func (p *SimpleProducer) declare() error {
	// declare exchange
	fmt.Printf("Binding exchange %v\n", p.exchange)
	if err := p.channel.ExchangeDeclare(
		p.exchange, // name of the exchange
		p.exchType, // type
		true,       // durable
		false,      // delete when complete
		false,      // internal
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("exchange declare error: %s", err)
	}
	return nil
}

// Close close producer
func (c *SimpleProducer) Close() error {
	return c.channel.Close()
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
