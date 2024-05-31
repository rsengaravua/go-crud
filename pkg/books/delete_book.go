package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	// Execute raw PostgreSQL query to delete the book
	result, err := h.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}
