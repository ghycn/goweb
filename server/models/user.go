package models

type User struct {
	BasicEntity
	UserName string `gorm:"column:username" json:"username"`
	Gender   int    `gorm:"column:gender" json:"gender"`
}
