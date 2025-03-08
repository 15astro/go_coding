package main

import (
 "fmt"
 "os"
)

type BankAccount struct{
	Name string
	Balance int64
}

func NewBankAccont(name string, balance int64) (*BankAccount, error){
 if balance <= 100{
	return nil, fmt.Errorf("Minimum balace required is 1000")
 }
 
 return &BankAccount{
	Name: name,
	Balance: balance,
 }, nil 
}

func (b *BankAccount) Deposit(amount int64){
	b.Balance += amount
	fmt.Println(b.Balance)
}

func main(){
    newAccount, err := NewBankAccont("Atish",2000)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(newAccount)
	newAccount.Deposit(1000)
	fmt.Println(newAccount.Balance)
	//fmt.Println(NewBankAccont("Atish",2000)
}

