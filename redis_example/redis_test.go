package main

import "testing"

func BenchmarkInitRedis(b *testing.B) {
	for i := 0; i < b.N; i ++{
		initRedis()
	}
}

func TestInitRedis(t *testing.T) {
	initRedis()
	for i := 0; i < 1000000; i ++ {
		main()
	}
}