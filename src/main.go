package main

import (
	"bank-account-test-task/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitializeRoutes(r)
	r.Run()
}
