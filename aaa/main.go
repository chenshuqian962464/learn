package awesomeProject

import (
	"github.com/chenshuqian962464/learn/aaa/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/index", handler.Index)
	router.POST("/user/create", handler.CreateUser)

	//router.Run(":8080")
}
