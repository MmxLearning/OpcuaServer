package service

import (
	"github.com/MmxLearning/OpcuaServer/internal/db/dao"
	"gorm.io/gorm"
)

var User = UserSrv{DB: dao.DB}

type UserSrv struct {
	*gorm.DB
}

func (a UserSrv) Begin() (UserSrv, error) {
	a.DB = a.DB.Begin()
	return a, a.Error
}

func (a UserSrv) Take(username string) (*dao.User, error) {
	model := dao.User{
		Username: username,
	}
	return &model, model.TakeByUsername(a.DB)
}
