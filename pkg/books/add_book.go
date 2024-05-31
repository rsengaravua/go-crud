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

	// Define the SQL query
	query := "INSERT INTO books (id, title, author, description) VALUES ($1, $2, $3, $4) RETURNING id"

	// Execute the query
	var id int
	err := h.DB.QueryRow(query, book.Id, book.Title, book.Author, book.Description).Scan(&id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// Update the book model with the returned ID
	book.Id = id

	// Respond with the created book
	ctx.JSON(http.StatusCreated, &book)
}
