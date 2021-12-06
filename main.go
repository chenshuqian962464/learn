package main

import (
	"github.com/chenshuqian962464/learn2/app"
	"github.com/chenshuqian962464/learn2/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	serverOpts := app.InitServerOpts()
	app.GlobalServerOpts = serverOpts
	router := gin.Default()

	router.GET("/index", handler.Index)
	router.POST("/user/create", handler.CreateUser)

	router.Run(":8080")
}
