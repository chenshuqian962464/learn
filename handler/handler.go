package handler

import (
	"github.com/chenshuqian962464/learn2/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(200, "ok")
	return
}

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.String(200, "name is empty!")
		return
	}
	age := c.PostForm("age")
	if len(age) == 0 {
		c.String(200, "age is empty")
		return
	}
	//fmt.Println(name,age)
	err := service.CreateUser(name, age)
	if err != nil {
		c.String(200, "has error :%v", err) //%v 自适应占位符
		return
	}
	c.String(200, "ok")
	return
}
