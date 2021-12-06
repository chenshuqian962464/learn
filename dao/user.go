package dao

import (
	"github.com/chenshuqian962464/learn2/app"
	"github.com/chenshuqian962464/learn2/model"
	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(u *model.User) error {
	//db,err  := sqlx.Connect("mysql","root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil{
	//	return err
	//}
	db := app.GlobalServerOpts.DB
	_, err := db.Exec("insert into user (firstname,age,high) values (?,?,?)", u.Firstname, u.Age, u.High)
	if err != nil {
		return err
	}
	return nil
}

//func QueryUser(){
//	db := app.GlobalServerOpts.DB
//	db.Select(nil,"select * from user where id=3")
//}
