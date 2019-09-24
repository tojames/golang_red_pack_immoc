package main

import (
	"testing"
)

func BenchmarkRunTimers(b *testing.B) {
	for _, count := range []int{1000, 2000, 5000, 10000, 20000, 50000, 100000, 500000} {
		RunTimers(count)
	}
	//for i := 0; i < b.N; i ++ {
	//	RunTimers(i)
	//}
}

