package dao

import "gorm.io/gorm"

type Opcua struct {
	ID        uint  `gorm:"primarykey"`
	CreatedAt int64 `gorm:"index"`

	Name      string `gorm:"not null;index;type:varchar(30)"`
	NodeID    string `gorm:"not null;index;type:varchar(255)"`
	Timestamp uint32 `gorm:"not null;index"`

	Data string
}

func (a *Opcua) Insert(tx *gorm.DB) error {
	return tx.Create(a).Error
}
