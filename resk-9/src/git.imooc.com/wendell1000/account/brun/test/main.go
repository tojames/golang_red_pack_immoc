package main

import (
	"encoding/json"
	"fmt"
	"git.imooc.com/wendell1000/account/services"
)

func main() {

	d, e := json.Marshal(&services.AccountTransferDTO{})
	fmt.Println(e)
	fmt.Println(string(d))
}
