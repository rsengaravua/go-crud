package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-crud/pkg/common/models"
)

type AddBookRequestBody struct {
	Id          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// Result holds the result of the database operation
type Result struct {
	Book  models.Book
	Error error
}

func (h *handler) AddBook(ctx *gin.Context) {
	var bookRequest AddBookRequestBody

	// Bind and validate request
	if err := ctx.ShouldBindJSON(&bookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a book model instance
	book := models.Book{
		Id:          bookRequest.Id,
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
	}

	// Channel to receive the result of the goroutine
	resultChan := make(chan Result)

	// Run the database operation in a goroutine
	go func() {
		if result := h.DB.Create(&book); result.Error != nil {
			resultChan <- Result{Book: book, Error: result.Error}
			return
		}
		resultChan <- Result{Book: book, Error: nil}
	}()

	// Receive the result from the goroutine
	result := <-resultChan

	// Handle the result
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// Respond with the created book
	ctx.JSON(http.StatusCreated, &result.Book)
}
