package controllers

import (
	"RestLecture/forms"
	"RestLecture/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FindBooks GET /books
func FindBooks(c *gin.Context) {
	var books []models.Book
	var requestParams forms.GetBooksParam

	c.Bind(&requestParams)
	fmt.Println(1111)
	fmt.Println(requestParams.Title)

	if requestParams.IsDesc {
		models.DB.Order("cost DESC").Where("is_published = ? and title LIKE ?", requestParams.Published, "%"+requestParams.Title+"%").Find(&books)
	} else {
		models.DB.Order("cost").Where("is_published = ? and title LIKE ?", requestParams.Published, "%"+requestParams.Title+"%").Find(&books)
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBook GET /books/:id
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	resp := map[string]any{
		"title":       book.Title,
		"description": book.Description,
		"cost":        book.Cost,
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// CreateBook POST /books
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{
		Title:       input.Title,
		Author:      input.Author,
		Cost:        input.Cost,
		Description: input.Description,
	}

	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook PUT /books/:id
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook DELETE /books/:id
func DeleteBook(c *gin.Context) {
	// Get model if exists
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

//func RateBook(c *gin.Context) {
//	// rate book if exists
//	var book models.Book
//	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
//		return
//	}
//}

func PublishBook(c *gin.Context) {
	// publish book if exists
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Model(&book).Updates(map[string]interface{}{
		"is_published": true,
	})
}

func UnPublishBook(c *gin.Context) {
	// unpublish book if exists
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Model(&book).Updates(map[string]interface{}{
		"is_published": false,
	})
}
