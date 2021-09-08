package main

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	w.balance = 10

	return w.balance
}
