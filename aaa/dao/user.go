package dao

import (
	"github.com/chenshuqian962464/learn/aaa/model"
	"github.com/jmoiron/sqlx"
)

func InsertUser(u *model.User) error {
	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	_, err = db.Exec("insert into user (firstname,age,high) values (?,?,?)", u.Firstname, u.Age, u.High)
	if err != nil {
		return err
	}
	return nil
}
