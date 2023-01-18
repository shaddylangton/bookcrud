package controllers

import (
	// "bookCRUD/models"  uftj-xagb-226k
	"projects/bookcrud/models"
	"net/http"
	_ "time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateBookInput struct {
	Booktittle    string `json:"book_tittle"`
	Bookcategory  string `json:"book_category"`
	BookNo        string `json:"book_no"`
	ISBNNo        string `json:"isbn_no"`
	PublisherName string `json:"publisher_name"`
	AuthorName    string `json:"author_name"`
	ShelfNo       uint   `json:"shelf_no"`
	RackNo        uint   `json:"rack_no"`
	Quantity      uint   `json:"quantity"`
	BookPrice     uint   `json:"book_price"`
	Description   string `json:"description"`
	Hireprice     uint   `json:"hire_price"`
}

type UpdateBookInput struct {
	Booktittle    string `json:"book_tittle"`
	Bookcategory  string `json:"book_category"`
	BookNo        string `json:"book_no"`
	ISBNNo        string `json:"isbn_no"`
	PublisherName string `json:"publisher_name"`
	AuthorName    string `json:"author_name"`
	ShelfNo       uint   `json:"shelf_no"`
	RackNo        uint   `json:"rack_no"`
	Quantity      uint   `json:"quantity"`
	BookPrice     uint   `json:"book_price"`
	Description   string `json:"description"`
	Hireprice     uint   `json:"hire_price"`
}

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new book

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	// Create task
	book := models.Book{
		Booktittle:    input.Booktittle,
		Bookcategory:  input.Bookcategory,
		BookNo:        input.BookNo,
		ISBNNo:        input.ISBNNo,
		PublisherName: input.PublisherName,
		AuthorName:    input.AuthorName,
		ShelfNo:       input.ShelfNo,
		RackNo:        input.RackNo,
		Quantity:      input.Quantity,
		BookPrice:     input.BookPrice,
		Description:   input.Description,
		Hireprice:     input.Hireprice,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput models.Book

	updatedInput.Booktittle = input.Booktittle
	updatedInput.Bookcategory = input.Bookcategory
	updatedInput.BookNo = input.BookNo
	updatedInput.ISBNNo = input.ISBNNo
	updatedInput.PublisherName = input.PublisherName
	updatedInput.AuthorName = input.AuthorName
	updatedInput.ShelfNo = input.ShelfNo
	updatedInput.RackNo = input.RackNo
	updatedInput.Quantity = input.Quantity
	updatedInput.BookPrice = input.BookPrice
	updatedInput.Description = input.Description
	updatedInput.Hireprice = input.Hireprice

	db.Model(&book).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
