package main

import "3-RabbitMQ/RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQRouting("exImooc","imooc_one")
	imoocOne.RecieveRouting()
}
