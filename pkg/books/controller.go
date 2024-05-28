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

	routes := router.Group("/books")
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.POST("/", h.AddBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
