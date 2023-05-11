package models

type User struct {
	ID       uint   `json:"id" gorm:"primary key"`
	Username string `json:"username" gorm:"unique; not null"`
	Password string `json:"password" gorm:"not null"`
}

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
