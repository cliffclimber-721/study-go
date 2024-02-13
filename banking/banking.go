package banking

type Account struct {
	Owner   string
	Balance int
}

type anoAccount struct {
	owner string
	bal   int
}

func NewAccount(owner string) *anoAccount {
	accs := &anoAccount{owner: owner, bal: 0}
	return &accs
}
