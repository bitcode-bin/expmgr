package main

type Wallet struct {
	balance int
}

func NewWallet(startingBalance int) *Wallet {
	return &Wallet{balance: startingBalance}
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) AddExpense(expense int) {
	w.balance = w.balance - expense
}

func (w *Wallet) AddIncome(income int) {
	w.balance = w.balance + income
}
