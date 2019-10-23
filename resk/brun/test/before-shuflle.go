package main

import (
	"fmt"
	"git.imooc.com/wendell1000/infra/algo"
)

func main() {
	count, amount := int64(10), int64(100)
	for i := int64(0); i < count; i++ {
		x := algo.BeforeShuffle(count, amount*100)
		fmt.Print(x, ",")

	}
	fmt.Println()
}
