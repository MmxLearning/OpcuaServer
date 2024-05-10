package controllers

import (
	"errors"
	"github.com/MmxLearning/OpcuaServer/internal/api/callback"
	"github.com/MmxLearning/OpcuaServer/internal/global"
	"github.com/MmxLearning/OpcuaServer/internal/pkg/jwt"
	"github.com/MmxLearning/OpcuaServer/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	"unsafe"
)

func Login(c *gin.Context) {
	var f struct {
		Username string `json:"username" form:"username" binding:"required,max=15"`
		Password string `json:"password" form:"password" binding:"required,max=100"`
	}
	if err := c.ShouldBind(&f); err != nil {
		callback.Error(c, callback.ErrForm, err)
		return
	}

	user, err := service.User.Take(f.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			time.Sleep(time.Second / 2)
			callback.Error(c, callback.ErrUsernameOrPasswordNotCorrect)
			return
		}
		callback.Error(c, callback.ErrDBOperation, err)
		return
	}

	if err = bcrypt.CompareHashAndPassword(
		unsafe.Slice(unsafe.StringData(user.Password), len(user.Password)),
		unsafe.Slice(unsafe.StringData(f.Password), len(f.Password)),
	); err != nil {
		callback.Error(c, callback.ErrUsernameOrPasswordNotCorrect, err)
		return
	}

	token, err := jwt.New(unsafe.Slice(unsafe.StringData(global.Config.JwtKey), len(global.Config.JwtKey))).
		GenerateToken(user.ID, time.Hour*24*30)
	if err != nil {
		callback.Error(c, callback.ErrUnexpected, err)
		return
	}

	callback.Success(c, gin.H{
		"token": token,
	})
}
