package banking

import (
	"errors"
	"fmt"
)

// 이런식으로 main.go 파일에서 에러를 따로 빼서 선언할 수 있는데
// if err != nil {
//		log.Fatalln(err)
//	}
// 따로 여기서 변수를 선언해서 에러를 표시할 수 있다.
// 앞에 무조건 err 붙여줘야 얘가 에러라는 걸 인식한다.
var errNoMoney = errors.New("NOOO MONEY!!")

type Account struct {
	Owner   string
	Balance int
}

type AnoAccount struct {
	own  string
	bals int
}

func NewAccount(owner string) *AnoAccount {
	accs := AnoAccount{own: owner, bals: 1000}
	return &accs
}

// a AnoAccount를 receiver 라고 부른다.
// Account2에 있는 변수를 사용할건데 이걸 이제 a라고 부르면서 쓸 계획
// import numpy as np 이면 np 쓴다는 말이랑 같은거다
func (a *AnoAccount) Deposit(amount int) {
	fmt.Println("DEPOSIT", amount)
	a.bals += amount
}

func (a *AnoAccount) Balances() int {
	return a.bals
}

// 실제로 이렇게 쓰면 결과값은 DEPOSIT 10 // 1000 이 나온다
// 내가 원하는 결과값은 1010인데 이렇게 값이 더해져서 나오지 않은 이유는
// 내가 지정한 a AnoAccount 는 AnoAccount에 있는 값을 받아온게 아니라 그냥 a 에 대한 변수값을 복사해서 갖다쓰는거나 다름없다.
// func를 통해서 지정한 함수를 쓰려면 포인터를 사용해서 그 안에 있는 함수를 호출해줘야한다.
// 포인터를 쓰게 되면 DEPOSIT 10 // 1010 이 나온다!

func (a *AnoAccount) Withdraw(amount int) error {
	if a.bals < amount {
		return errNoMoney
	}
	a.bals -= amount
	// nil은 None null이랑 같은 뜻이다.
	return nil
}

// 흔히 우리가 아는 은행 계좌에는 마이너스 통장이 아니라면 0에서 끝내야하는 게 맞다.
// 0 이하가 되면 오류가 뜨도록 error handling을 해줘야한다.
// Go 언어에선 Python 처럼 try-except문이 없기 때문에 if-else로 조건을 걸어준다.

//func forNewAccountMain() {
//	accountCho := banking.Account{Owner: "chocho", Balance: 1000}
//	accountShin := banking.Account{Owner: "shin"}
//	acc := banking.NewAccount("kangs")
//	fmt.Println(accountCho)
//	fmt.Println(accountShin)
//	fmt.Println(acc)
//}

// 밑에 main 함수 사용할거면 import log 해야한다!!

//func WithdrawAndDepositmain() {
//	acc := banking.NewAccount("chocho")
//	acc.Deposit(10)
//	fmt.Println(acc.Balances())
// error handling
//	err := acc.Withdraw(2000)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(acc.Balances())
//}
