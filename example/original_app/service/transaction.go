package service

type TransactionService struct {
}

func (t TransactionService) Id() string {
	return "transactions"
}

func (t TransactionService) Provide(userId uint64) interface{} {
	return []interface{}{
		map[string]interface{}{
			"id":       1,
			"type":     "pay",
			"amount":   96623,
			"currency": "CNY",
		},
		map[string]interface{}{
			"id":       2,
			"type":     "repay",
			"amount":   10000,
			"currency": "CNY",
		},
	}
}

func (t TransactionService) StrongDepend() bool {
	return true
}
