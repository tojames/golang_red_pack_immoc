package main

import (
	"fmt"
	"sort"
)

type uints []uint

func NewUints() *uints {
	return &uints{500, 130, 160, 30, 560, 910, 1, 3, 9, 7, 6}
}

func (ui uints) Len() int {
	return len(ui)
}

func (ui uints) Less(i, j int) bool {
	return ui[i] < ui[j]
}

func (ui uints) Swap(i, j int) {
	ui[i], ui[j] = ui[j], ui[i]
}

func (ui uints) search(key uint) int {
	//查找算法
	f := func(x int) bool {
		return ui[x] > key
	}

	i := sort.Search(len(ui), f)
	return i
}

func main() {
	ui := NewUints()
	sort.Sort(ui)
	fmt.Printf("ui type is %v\n", ui)
	rs := ui.search(1)
	fmt.Printf("sort result is %v\n", rs)
}
