package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"net/http"
)

const (
	AccNotCreatedErr = "Account is not created"
	AccisClosedErr   = "Account is closed"
)

type Account struct {
	balance int64
	mutex   sync.Mutex
	isOpen  bool
}

var (
	router *gin.Engine
)

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run()
}

func initializeRoutes() {
	router.POST("account", PostAccount)
	router.GET("account", GetAccount)
	router.PUT("account", PutAccount)
	router.DELETE("account", DeleteAccount)
}

func PostAccount(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}

func GetAccount(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}

func PutAccount(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}

func DeleteAccount(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}

// Open creates new pointer to Account object with a given initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{isOpen: true, balance: initialDeposit, mutex: sync.Mutex{}}
}

// Close closes bank account and returns current balance
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.isOpen {
		return 0, false
	}
	curBalance := a.balance
	a.isOpen = false
	a.balance = 0
	return curBalance, true
}

// Balance returns account balance
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.isOpen {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.isOpen || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
