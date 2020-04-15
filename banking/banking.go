package banking

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	Owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

func NewAccount(owner string) *Account {
	account := Account{Owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("Deposit Amount", amount)
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

//Withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
