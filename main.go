package main

import (
	"fmt"

	"github.com/cliffclimber-721/study-go/banking"
)

func main() {
	accountCho := banking.Account{Owner: "chocho", Balance: 1000}
	accountShin := banking.Account{Owner: "shin"}
	acc := banking.NewAccount("kangs")
	fmt.Println(accountCho)
	fmt.Println(accountShin)
	fmt.Println(acc)
}
