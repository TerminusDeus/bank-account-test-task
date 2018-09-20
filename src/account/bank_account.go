package account

import "sync"

type Account struct {
	balance int64
	mutex   sync.Mutex
	IsOpen  bool
}

// Open creates new pointer to Account object with a given initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{IsOpen: true, balance: initialDeposit, mutex: sync.Mutex{}}
}

// Close closes bank account and returns current balance
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.IsOpen {
		return 0, false
	}
	curBalance := a.balance
	a.IsOpen = false
	a.balance = 0
	return curBalance, true
}

// Balance returns account balance value
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.IsOpen {
		return 0, false
	}
	return a.balance, true
}

// Deposit changes account balance value
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.IsOpen || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}


