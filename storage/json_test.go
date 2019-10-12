package storage

import (
	"fmt"
	"testing"
)

func isNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSONStorage(t *testing.T) {
	storage, err := NewJSONStorage("testdata.json")
	isNil(t, err)

	defer storage.Close()

	fmt.Println(storage.Balance())
	//storage.SetBalance(100)

	storage.NewTransaction(50)
	storage.NewTransaction(-20)
}
