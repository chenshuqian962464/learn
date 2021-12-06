package service

import (
	"github.com/chenshuqian962464/learn/aaa/dao"
	"github.com/chenshuqian962464/learn/aaa/model"
	"strconv"
)

func CreateUser(name, age string) error {
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return err
	}
	user := &model.User{ //ID 不用写 因为数据库设置了自增
		Firstname: name,
		Age:       ageInt,
		High:      180,
	}
	err = dao.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}
