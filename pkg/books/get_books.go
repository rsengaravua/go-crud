package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-crud/pkg/common/models"
)

func (h *handler) GetBooks(ctx *gin.Context) {
	rows, err := h.DB.Query("SELECT id, title, author, description FROM books")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Description); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}
