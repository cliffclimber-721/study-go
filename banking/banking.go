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
