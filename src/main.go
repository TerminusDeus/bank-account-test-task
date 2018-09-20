package main

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	acc "bank-account-test-task/src/account"
)

type PutAccountReq struct {
	InitialAmount int64 `json:"initialAmount" form:"initialAmount"`
}

type PostAccountReq struct {
	Amount int64 `json:"amount" form:"amount"`
}

type GetAccountResp struct {
	Amount int64 `json:"amount"`
}

const (
	AccIsNotCreatedErr = "Account is not created"
	AccIsClosedErr     = "Account is closed"
	NotEnoughMoney     = "Not enough money"
)

var (
	router  *gin.Engine
	account *acc.Account
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

//  PostAccount lets change account balance value
func PostAccount(context *gin.Context) {
	var postAccountReq PostAccountReq
	var amount int64

	if err := context.Bind(&postAccountReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Context bind error: ", err.Error())})
		return
	} else {
		amount = postAccountReq.Amount
	}

	if amount == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Amount value cannot be zero")})
		return
	}

	if account == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": AccIsNotCreatedErr})
		return
	} else {
		if !account.IsOpen {
			context.JSON(http.StatusBadRequest, gin.H{"error": AccIsClosedErr})
			return
		} else {
			if _, ok := account.Deposit(amount); !ok {
				context.JSON(http.StatusBadRequest, gin.H{"error": NotEnoughMoney})
				return
			}
		}
	}

	context.JSON(http.StatusOK, "Account successfully created")
}

//  GetAccount lets getting account balance
func GetAccount(context *gin.Context) {
	var amount int64
	var ok bool
	if account == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": AccIsNotCreatedErr})
		return
	} else {
		if !account.IsOpen {
			context.JSON(http.StatusBadRequest, gin.H{"error": AccIsClosedErr})
			return
		} else {
			if amount, ok = account.Balance(); !ok {
				context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Account balance value cannot be defined")})
				return
			}
		}
	}
	context.JSON(http.StatusOK, GetAccountResp{amount})
}

// PutAccount lets new account creation
func PutAccount(context *gin.Context) {
	var putAccountReq PutAccountReq
	var initialAmount int64

	if err := context.Bind(&putAccountReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Context bind error: ", err.Error())})
		return
	} else {
		initialAmount = putAccountReq.InitialAmount
	}

	if initialAmount < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Initial deposit cannot be negative")})
		return
	}

	if account == nil {
		account = acc.Open(initialAmount)
		if account == nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": AccIsNotCreatedErr})
			return
		}
	}

	context.JSON(http.StatusOK, "Account successfully created")
}

// DeleteAccount handles existing account closing logic
func DeleteAccount(context *gin.Context) {
	if account == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": AccIsNotCreatedErr})
		return
	}

	if _, ok := account.Close(); !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Account cannot be deleted"})
		return
	}

	context.JSON(http.StatusOK, "Account successfully deleted")
}
