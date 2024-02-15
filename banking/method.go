package banking

type Account2 struct {
	Owner   string
	Balance int
}

// a Account2를 receiver 라고 부른다.
// Account2에 있는 변수를 사용할건데 이걸 이제 a라고 부르면서 쓸 계획
// import numpy as np 이면 np 쓴다는 말이랑 같은거다
func (a Account2) Deposit(amount int) {
	a.Balance += amount

}
