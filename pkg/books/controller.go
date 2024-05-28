package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/v1")
	routes.GET("/AllBooks", h.GetBooks)
	routes.GET("/book/:id", h.GetBook)
	routes.POST("/book", h.AddBook)
	routes.PUT("/book/:id", h.UpdateBook)    // Corrected route definition for updating a book
	routes.DELETE("/book/:id", h.DeleteBook) // Corrected route definition for deleting a book
}
