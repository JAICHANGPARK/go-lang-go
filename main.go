package main

import (
	"fmt"

	"github.com/JAICHANGPARK/learngo/banking"
)

func main() {
	fmt.Println("Hello")
	acc := banking.NewAccount("Dream")
	fmt.Println(acc)
	acc.Deposit(10)
	fmt.Println(acc.Balance())
	acc.Withdraw(20)
	fmt.Println(acc.Balance())
}
