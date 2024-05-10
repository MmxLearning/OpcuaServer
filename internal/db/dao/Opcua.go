package dao

type Opcua struct {
	ID uint `gorm:"primarykey"`

	Name   string `gorm:"not null;index;type:varchar(30)"`
	NodeID string `gorm:"not null;index;type:varchar(255)"`

	Data string
}
