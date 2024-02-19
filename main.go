package main

import (
	"fmt"

	"github.com/cliffclimber-721/study-go/banking"
)

func main() {
	acc := banking.NewAccount("chocho")
	acc.Deposit(10)
	fmt.Println(acc)
}
