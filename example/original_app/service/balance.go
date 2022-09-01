package service

type BalanceService struct {
}

func (b BalanceService) Id() string {
	return "balance"
}

func (b BalanceService) Provide(userId uint64) interface{} {
	return 96.623
}

func (b BalanceService) StrongDepend() bool {
	return true
}
