package dao

import (
	"gorm.io/gorm"
)

type Opcua struct {
	ID        uint  `gorm:"primarykey" json:"id"`
	CreatedAt int64 `gorm:"index" json:"created_at"`

	Name      string `gorm:"not null;index;type:varchar(30)" json:"name"`
	NodeID    string `gorm:"not null;index;type:varchar(255)" json:"node_id"`
	Timestamp uint32 `gorm:"not null;index" json:"timestamp"`

	Data string `json:"data"`
}

func (a *Opcua) Insert(tx *gorm.DB) error {
	return tx.Create(a).Error
}

func (a *Opcua) Search(tx *gorm.DB, startTime, endTime int64) ([]Opcua, error) {
	var result = make([]Opcua, 0)
	tx = tx.Model(a).Where(a)
	if startTime != 0 {
		tx = tx.Where("timestamp>?", startTime)
	}
	if endTime != 0 {
		tx = tx.Where("timestamp<?", endTime)
	}
	return result, tx.Find(&result).Error
}
