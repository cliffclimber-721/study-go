package main

import (
	"fmt"
	"log"

	"github.com/cliffclimber-721/study-go/banking"
)

func main() {
	acc := banking.NewAccount("chocho")
	acc.Deposit(10)
	fmt.Println(acc.Balances())
	// error handling
	err := acc.Withdraw(2000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(acc.Balances())
}
