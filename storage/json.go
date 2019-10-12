package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

type transaction struct {
	Timestamp time.Time
	Amount    int
}

type account struct {
	Balance      int
	Transactions []*transaction
}

type JSONStorage struct {
	filename string
	data     *account
}

func NewJSONStorage(filename string) (*JSONStorage, error) {
	js := &JSONStorage{
		filename: filename,
		data:     &account{},
	}

	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if err = js.readFile(); err != nil {
			return nil, errors.Wrap(err, "read file")
		}
	}
	return js, nil
}

func (js *JSONStorage) Close() {
	js.writeFile()
}

func (js *JSONStorage) Balance() (int, error) {
	return js.data.Balance, nil
}

func (js *JSONStorage) SetBalance(balance int) error {
	js.data.Balance = balance
	return nil
}

func (js *JSONStorage) NewTransaction(amount int) error {
	trans := &transaction{
		Amount:    amount,
		Timestamp: time.Now(),
	}
	js.data.Balance += amount
	js.data.Transactions = append(js.data.Transactions, trans)
	return nil
}

func (js *JSONStorage) writeFile() error {
	jsonData, err := json.Marshal(js.data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(js.filename, jsonData, 0644)
}

func (js *JSONStorage) readFile() error {
	data, err := ioutil.ReadFile(js.filename)
	if err != nil {
		return err
	}

	var acc account
	err = json.Unmarshal(data, &acc)
	if err != nil {
		return err
	}

	js.data = &acc
	return nil
}
