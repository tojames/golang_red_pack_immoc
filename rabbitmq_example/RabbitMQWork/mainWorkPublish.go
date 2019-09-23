package main

import (
	"fmt"
	"3-RabbitMQ/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"imoocSimple")

	for i := 0; i <= 100; i++ {
		fmt.Println(strconv.Itoa(i))
		rabbitmq.PublishSimple("Hello imooc!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
