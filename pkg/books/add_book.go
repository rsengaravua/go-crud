package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsengaravua/go-crud/pkg/common/models"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h *handler) AddBook(ctx *gin.Context) {
	bookRequest := AddBookRequestBody{}

	if err := ctx.BindJSON(&bookRequest); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
	}

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
