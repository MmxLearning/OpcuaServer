package dao

import "gorm.io/gorm"

type Opcua struct {
	ID uint `gorm:"primarykey"`

	Name   string `gorm:"not null;index;type:varchar(30)"`
	NodeID string `gorm:"not null;index;type:varchar(255)"`

	Data string
}

func (a *Opcua) Insert(tx *gorm.DB) error {
	return tx.Create(a).Error
}
