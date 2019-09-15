package inmemory

import (
	"sync"
)

type Wallet struct {
	balance int
	mx      sync.RWMutex
}

func NewWallet(startingBalance int) *Wallet {
	return &Wallet{balance: startingBalance}
}

func (w *Wallet) Balance() int {
	w.mx.RLock()
	defer w.mx.RUnlock()
	return w.balance
}

func (w *Wallet) AddExpense(expense int) int {
	w.mx.Lock()
	defer w.mx.Unlock()
	w.balance = w.balance - expense
	return w.balance
}

func (w *Wallet) AddIncome(income int) int {
	w.mx.Lock()
	defer w.mx.Unlock()
	w.balance = w.balance + income
	return w.balance
}
