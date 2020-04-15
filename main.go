package main

import (
	"fmt"

	"github.com/JAICHANGPARK/learngo/mydict"
)

func main() {
	// fmt.Println("Hello")
	// acc := banking.NewAccount("Dream")
	// fmt.Println(acc)
	// acc.Deposit(10)
	// fmt.Println(acc.Balance())
	// err := acc.Withdraw(20)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(acc.Balance(), acc.Onwer())

	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "Hello"
	fmt.Println(dictionary)
}
