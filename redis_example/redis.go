package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	redisHost = "127.0.0.1:6379"
	redisDb = 1
	//password = "rmt2019!"
)

var redisClient *redis.Pool


func initRedis() {
	redisClient = &redis.Pool{
		MaxActive:   100,                              //  最大连接数，即最多的tcp连接数，一般建议往大的配置，但不要超过操作系统文件句柄个数（centos下可以ulimit -n查看）
		MaxIdle:     100,                              // 最大空闲连接数，即会有这么多个连接提前等待着，但过了超时时间也会关闭。
		IdleTimeout: time.Duration(100) * time.Second, // 空闲连接超时时间，但应该设置比redis服务器超时时间短。否则服务端超时了，客户端保持着连接也没用
		Wait:        true,                             // 当超过最大连接数 是报错还是等待， true 等待 false 报错
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisHost, redis.DialDatabase(redisDb), /*redis.DialPassword(password)*/)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func main() {
	initRedis()
	redisConn := redisClient.Get()
	defer redisConn.Close()
	rs, err := redisConn.Do("set", "test", "world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("type is %T, value is %v\n", rs, rs)
}