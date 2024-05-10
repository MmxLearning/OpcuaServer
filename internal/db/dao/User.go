package dao

import "gorm.io/gorm"

type User struct {
	ID uint `gorm:"primarykey"`

	Nickname string `gorm:"not null"`
	Username string `gorm:"type:varchar(15);uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (a *User) TakeByUsername(tx *gorm.DB) error {
	return tx.Model(a).Where(a, "username").Take(a).Error
}
