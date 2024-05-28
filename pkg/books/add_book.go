package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(ctx *gin.Context) {
	bookRequest := AddBookRequestBody{}

	if err := ctx.BindJSON(&bookRequest); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.book
	book.Title = bookRequest.Title
	book.Author = bookRequest.Author
	book.Description = bookRequest.Description

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
