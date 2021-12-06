package handler

import (
	"github.com/chenshuqian962464/learn/aaa/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(200, "you are right")
	return
}

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.String(200, "name is empty")
		return
	}
	age := c.PostForm("age")
	if len(age) == 0 {
		c.String(200, "age is empty")
		return
	}
	err := service.CreateUser(name, age)
	if err != nil {
		c.String(200, "has error :%v", err)
	}
	c.String(200, "you are right")
	return
}
