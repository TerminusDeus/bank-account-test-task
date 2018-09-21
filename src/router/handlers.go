package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	acc "bank-account-test-task/src/account"
)

var account *acc.Account

type PutAccountReq struct {
	InitialAmount int64 `json:"initialAmount,required"`
}

type PostAccountReq struct {
	Amount int64 `json:"amount,required"`
}

type GetAccountResp struct {
	Amount int64 `json:"amount"`
}

func InitializeRoutes(r *gin.Engine) {
	r.POST("account", PostAccountHandler)
	r.GET("account", GetAccountHandler)
	r.PUT("account", PutAccountHandler)
	r.DELETE("account", DeleteAccountHandler)
}

// PostAccountHandler lets change account balance value
func PostAccountHandler(context *gin.Context) {
	var postAccountReq PostAccountReq

	if err := context.Bind(&postAccountReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Context bind error: %s\n", err.Error())})
		return
	}

	if postAccountReq.Amount == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": AmValCantBe0Err})
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
			if _, ok := account.Deposit(postAccountReq.Amount); !ok {
				context.JSON(http.StatusBadRequest, gin.H{"error": NotEnoughMoneyErr})
				return
			}
		}
	}

	context.JSON(http.StatusOK, "Account balance successfully changed")
}

// GetAccountHandler lets getting account balance
func GetAccountHandler(context *gin.Context) {
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
				context.JSON(http.StatusBadRequest, gin.H{"error": AccValCantBeDefinedErr})
				return
			}
		}
	}
	context.JSON(http.StatusOK, GetAccountResp{amount})
}

// PutAccountHandler lets new account creation
func PutAccountHandler(context *gin.Context) {
	var putAccountReq PutAccountReq
	var initialAmount int64

	if err := context.Bind(&putAccountReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Context bind error: %s\n", err.Error())})
		return
	} else {
		initialAmount = putAccountReq.InitialAmount
	}

	if initialAmount < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": InitDepositCantBeNegErr})
		return
	}

	if account == nil {
		account = acc.Open(initialAmount)
		if account == nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": AccIsNotCreatedErr})
			return
		}
	} else {
		if !account.IsOpen {
			account.IsOpen = true
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": AccHasBeenAlreadyCreatedErr})
			return
		}
	}
	context.JSON(http.StatusOK, "Account successfully created")
}

// DeleteAccountHandler handles existing account closing logic
func DeleteAccountHandler(context *gin.Context) {
	if account == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": AccIsNotCreatedErr})
		return
	}

	if _, ok := account.Close(); !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": AccCantBeDeletedErr})
		return
	}

	context.JSON(http.StatusOK, "Account successfully deleted")
}
