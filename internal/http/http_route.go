package http

import (
    "github.com/gin-gonic/gin"
	"byfood_service/internal/handler"
)

func Route(r *gin.Engine, handler *handler.BookHandler) {
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    r.GET("/books", handler.GetBooks)
    r.GET("/book/:id", handler.GetBookByID)
    r.POST("/book", handler.CreateBook)
    r.PUT("/book/:id", handler.UpdateBook)
    r.DELETE("/book/:id", handler.DeleteBook)
}
