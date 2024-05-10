package dao

type User struct {
	ID uint `gorm:"primarykey"`

	Nickname string `gorm:"not null"`
	Username string `gorm:"type:varchar(15);uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
