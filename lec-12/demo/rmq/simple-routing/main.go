package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rmqroute/consumer"
	"rmqroute/pkg/driver/rabbitmq"
	"rmqroute/producer"
	"sync"
)

func main() {
	// creates rmq connection
	rmqURI := "amqp://guest:guest@localhost:5672/"
	exch := "order"
	exchType := "direct"
	queue := "order_processor"
	routingKey := ""

	// connect to rabbitmq
	rmq := rabbitmq.NewRMQ(rmqURI)
	// create 1 channel for producer
	pCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel: ", err)
		return
	}
	// create 1 channel for consumer
	cCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel: ", err)
		return
	}

	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg.Add(2) // add 2 for consumer and producer
	producer := producer.NewSimpleProducer(ctx, wg, pCh, exch, exchType, routingKey)
	consumer := consumer.NewSimpleConsumer(ctx, wg, cCh, exch, exchType, routingKey, queue)

	/////////////////////////////////////////////////
	// graceful shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c // waits for the termination signal
		fmt.Printf("Got %s signal. Exiting...\n", sig)
		producer.Close() // stop scheduler at the end
		consumer.Close()
		cancelFunc()
	}()
	/////////////////////////////////////////////////

	// Start now
	go consumer.Start()
	go producer.Start()

	// wait for exit signal Ctrl - C
	wg.Wait()
}
