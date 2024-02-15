package banking

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

type AnoAccount struct {
	own  string
	bals int
}

func NewAccount(owner string) *AnoAccount {
	accs := AnoAccount{own: owner, bals: 0}
	return &accs
}

func forMain() {
	accountCho := banking.Account{Owner: "chocho", Balance: 1000}
	accountShin := banking.Account{Owner: "shin"}
	acc := banking.NewAccount("kangs")
	fmt.Println(accountCho)
	fmt.Println(accountShin)
	fmt.Println(acc)
}
