package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"byfood_service/internal/domain"
	"byfood_service/internal/service"
)

type BookHandler struct {
    service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
    return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
    books, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
        return
    }
    c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
    id := c.Param("id")

    book, err := h.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
    var input struct {
        Title  string `json:"title" binding:"required"`
        Author string `json:"author" binding:"required"`
        Year   int    `json:"year" binding:"required,gte=0"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book := domain.Book{
        Title:  input.Title,
        Author: input.Author,
        Year:   input.Year,
    }

    if err := h.service.Create(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
        return
    }
    c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var input struct {
        Title  string `json:"title" binding:"required"`
        Author string `json:"author" binding:"required"`
        Year   int    `json:"year" binding:"required,gte=0"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedData := domain.Book{
        Title:  input.Title,
        Author: input.Author,
        Year:   input.Year,
    }

    book, err := h.service.Update(id, &updatedData)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or failed to update"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
    id := c.Param("id")
    err := h.service.Delete(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID or failed to delete"})
        return
    }
    c.Status(http.StatusNoContent)
}
