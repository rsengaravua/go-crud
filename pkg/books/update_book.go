package books

import (
	"database/sql"
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
	id := ctx.Param("id")

	var bookBody UpdateBookRequestBody
	if err := ctx.ShouldBindJSON(&bookBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Execute raw PostgreSQL query to update the book
	result, err := h.DB.Exec("UPDATE books SET title = $1, author = $2, description = $3 WHERE id = $4", bookBody.Title, bookBody.Author, bookBody.Description, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Fetch the updated book details
	var book models.Book
	err = h.DB.QueryRow("SELECT id, title, author, description FROM books WHERE id = $1", id).Scan(&book.Id, &book.Title, &book.Author, &book.Description)
	if err == sql.ErrNoRows {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
