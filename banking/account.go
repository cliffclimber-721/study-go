package banking

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

// a AnoAccount를 receiver 라고 부른다.
// Account2에 있는 변수를 사용할건데 이걸 이제 a라고 부르면서 쓸 계획
// import numpy as np 이면 np 쓴다는 말이랑 같은거다
func (a AnoAccount) Deposit(amount int) {
	a.bals += amount

}

//func forMain() {
//	accountCho := banking.Account{Owner: "chocho", Balance: 1000}
//	accountShin := banking.Account{Owner: "shin"}
//	acc := banking.NewAccount("kangs")
//	fmt.Println(accountCho)
//	fmt.Println(accountShin)
//	fmt.Println(acc)
//}
