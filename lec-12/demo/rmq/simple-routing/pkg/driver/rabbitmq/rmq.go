package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RMQ struct {
	URI        string `json:"uri"`
	connection *amqp.Connection
}

func NewRMQ(uri string) *RMQ {
	connection, err := amqp.Dial(uri)
	if err != nil {
		fmt.Println("Cannot connect to rabbitmq: ", uri, err)
		return nil
	}
	return &RMQ{
		URI:        uri,
		connection: connection,
	}
}

func (rmq *RMQ) GetChannel() (*amqp.Channel, error) {
	return rmq.connection.Channel()
}

func (rmq *RMQ) Close() {
	rmq.connection.Close()
}
