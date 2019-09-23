package main

import "3-RabbitMQ/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.ConsumeSimple()
}
