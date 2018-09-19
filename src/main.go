package src

import (
	"github.com/gin-gonic/gin"
	"sync"
	"bank-account-test-task/src/common"
	"net/http"
)

const (
	AccNotCreatedErr = "Account is not created"
	AccisClosedErr   = "Account is closed"
)

type Account struct {
	balance float64
	mutex   sync.Mutex
	isOpen  bool
}

var (
	Config common.Config
	router *gin.Engine
)

func main() {
	Config = common.InitConfig()
	router = gin.Default()
	initializeRoutes()

	// port can be customized in conf.json
	router.Run(":" + Config.Port)
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

// Open creates new Account pointer object with a given initial amount
func Open(initialAmount float64) *Account {
	if initialAmount < 0 {
		return nil
	}
	return &Account{isOpen: true, balance: initialAmount, mutex: sync.Mutex{}}
}

// Close closes bank account and returns current balance
func (a *Account) Close() (out float64, ok bool) {
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
func (a *Account) Balance() (balance float64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.isOpen {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount float64) (newBalance float64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.isOpen || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
