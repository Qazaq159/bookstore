// models/book.go

package models

type Book struct {
	ID          uint    `json:"id" gorm:"primary key"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type GetBook struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type CreateBookInput struct {
	Title       string  `json:"title"  binding:"required"`
	Author      string  `json:"author" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Cost        float64 `json:"cost" binding:"required"`
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
