package main

type Wallet struct {
	balance int
}

type WalletService interface {
	Balance() int
	AddExpense(int) int
	AddIncome(int) int
}
