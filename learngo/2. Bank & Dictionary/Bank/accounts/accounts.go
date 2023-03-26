package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// Creates new account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on the account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Get balance of the account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount on the account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("not enough balance")
	}
	a.balance -= amount
	return nil
}

// Change owner on the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Get owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account.\nBalance: ", a.balance)
}
