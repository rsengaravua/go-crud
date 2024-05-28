package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-crud/pkg/common/models"
)

type UpdateBookRequestBody struct {
	// Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h *handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")

	bookBody := UpdateBookRequestBody{}

	if err := ctx.BindJSON(&bookBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&bookBody, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = bookBody.Title
	book.Author = bookBody.Author
	book.Description = bookBody.Description

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}
