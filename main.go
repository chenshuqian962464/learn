package main

import (
	"fmt"
	//"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int64  `db:"id" json:"id,omitempty"`
	Firstname string `db:"firstname" json:"firstname,omitempty"`
	Age       int    `db:"age" json:"age,omitempty"`
	High      int    `db:"high" json:"high,omitempty"`
}

//type User struct {
//	ID        int64  `db:"id" json:"id,omitempty"`
//	Firstname string `db:"firstname" json:"firstname,omitempty"`
//	Age       int    `db:"age" json:"age,omitempty"`
//	High      int    `db:"high" json:"high,omitempty"`
//}

func main() {
	router := gin.Default()
	router.POST("/user/create", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		//fmt.Println(name, age)
		DBClient, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			c.String(200, err.Error())
			return
		}
		_, err = DBClient.Exec(`insert into user(firstname,age,high)values(?,?,180)`, name, age)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		c.String(200, "you are right")
	})
	router.GET("/user/:ID", func(c *gin.Context) {
		ID := c.Param("ID")
		fmt.Println(ID)
		DB, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			c.String(200, "err.Error()")
			return
		}
		user := new(User)
		err = DB.Get(user, "select * from user where id=?", ID)
		if err != nil {
			c.String(200, "err.Error()")
			return
		}
		c.JSON(200, user)

	})
	//router.Run(":8080")
}

////import (
////	"fmt"
////	"github.com/gin-gonic/gin"
////	_ "github.com/go-sql-driver/mysql"
////	"github.com/jmoiron/sqlx"
////)
////
////type User struct {
////	ID        int64  `db:"id" json:"id,omitempty"`
////	Firstname string `db:"firstname" json:"firstname,omitempty"`
////	Age       int    `db:"age" json:"age,omitempty"`
////	High      int    `db:"high" json:"high,omitempty"`
////}
////
////func main() {
////	engine := gin.Default()
////	engine.POST("/user/create", func(c *gin.Context) {
////		name := c.PostForm("name")
////		age := c.PostForm("age")
////		//fmt.Println(name, age)
////		DB, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
////		if err != nil {
////			c.String(200, err.Error())
////			return
////		}
////		_, err = DB.Exec(`insert into user(firstname,age,high) values (?,?,180)`, name, age)
////		if err != nil {
////			c.String(200, err.Error())
////			return
////		}
////		c.String(200, "ok")
////	})
////	engine.GET("/index", func(c *gin.Context) {
////		//c.String(200,"<h1>come on </h1>")
////		c.Header("Content-Type", "text/html; charset=utf-8")
////		c.String(200, "<h1>aaaa</h1>")
////	})
////	engine.GET("/user/:ID", func(c *gin.Context) {
////		ID := c.Param("ID")
////		fmt.Println(ID)
////		DB, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
////		if err != nil {
////			c.String(200, err.Error())
////			return
////		}
////		user := new(User)
////		err = DB.Get(user, "select * from user where id=?", ID)
////		if err != nil {
////			c.String(200, err.Error())
////			return
////		}
////		c.JSON(200, user)
////
////	})
////
////	engine.Run(":8080")
////}
