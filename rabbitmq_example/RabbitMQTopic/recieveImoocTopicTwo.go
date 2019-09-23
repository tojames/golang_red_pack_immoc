package main

import "3-RabbitMQ/RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQTopic("exImoocTopic","imooc.*.two")
	imoocOne.RecieveTopic()
}
