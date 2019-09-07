package main

import "testing"

func TestWallet(t *testing.T) {
	w := NewWallet(300)
	var bal, want int

	bal = w.Balance()
	if bal != 300 {
		t.Fatalf("want=%d, got=%d", 300, bal)
	}

	w.AddIncome(300)
	want = 600
	bal = w.Balance()
	if bal != want {
		t.Fatalf("want=%d, got=%d", want, bal)
	}

	w.AddExpense(200)
	want = 400
	bal = w.Balance()
	if bal != want {
		t.Fatalf("want=%d, got=%d", want, bal)
	}
}
