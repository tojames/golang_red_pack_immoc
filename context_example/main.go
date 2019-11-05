package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	go handle(ctx, 100 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <- ctx.Done():
		fmt.Println("handle ", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func inc(a int) int {
	res := a + 1
	time.Sleep(1 * time.Second)
	return res
}

func Add(ctx context.Context, a, b int) (res int) {
	for i := 0; i < a; i ++ {
		res = inc(res)
		select {
		case <- ctx.Done():
			return -1
		default:

		}
	}

	for i := 0; i < b; i ++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}
	return
}

func main1() {
	{
		a := 1
		b := 2
		timeout := 2 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)

		res := Add(ctx, 1, 2)
		fmt.Printf("Compute: %d + %d, result: %d\n", a, b, res)
	}

	{
		// 手动取消
		a := 1
		b := 2
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancel() // 在调用处主动取消
		}()
		res := Add(ctx, 1, 2)
		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
}
