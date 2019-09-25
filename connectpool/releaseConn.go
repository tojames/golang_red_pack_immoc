package main

import (
	"context"
	"golang/connectpool/pool"
)

func main() {

	ctx := context.Background()
	config := &pool.Config{
		MaxConn: 2,
		MaxIdle: 1,
	}
	conn := pool.Prepare(ctx, config)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	if _, err := conn.New(ctx); err != nil {
		return
	}
	if _, err := conn.New(ctx); err != nil {
		return
	}
}