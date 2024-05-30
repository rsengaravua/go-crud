package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-crud/pkg/common/models"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h *handler) UpdateBook(ctx *gin.Context) {
	idStr := ctx.Param("id")

	bookBody := new(UpdateBookRequestBody)

	if err := ctx.BindJSON(&bookBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	// Fetch existing book
	if result := h.DB.First(&book, "id = ?", idStr); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// Update book attributes
	book.Title = bookBody.Title
	book.Author = bookBody.Author
	book.Description = bookBody.Description

	// Execute raw PostgreSQL query to update the book
	if err := h.DB.Exec("UPDATE books SET title=?, author=?, description=? WHERE id=?", book.Title, book.Author, book.Description, idStr).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
