package main

type DBService interface {
	Balance() (int, error)
	SetBalance(int) error
	NewTransaction(int) error
}
