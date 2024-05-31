package books

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-crud/pkg/common/models"
)

func (h *handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	// Execute raw PostgreSQL query to get the book
	err := h.DB.QueryRow("SELECT id, title, author, description FROM books WHERE id = $1", id).Scan(&book.Id, &book.Title, &book.Author, &book.Description)
	if err == sql.ErrNoRows {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
