package service

import (
	"github.com/chenshuqian962464/learn2/dao"
	"github.com/chenshuqian962464/learn2/model"
	"strconv"
)

func CreateUser(name, age string) error {
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return err
	}

	user := &model.User{ //因为ID数据库设置的是自增  所以不用写
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
