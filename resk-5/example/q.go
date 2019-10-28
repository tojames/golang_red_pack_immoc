package main

import (
	"fmt"
	"path"
)

func main() {
	p := "/Users/tietang/my/gitcode/resk-projects/src/imooc.com/resk/infra/base/dbx.go"
	d, f := path.Split(p)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(path.Base(d))
	fmt.Println(path.Base(p))
}
