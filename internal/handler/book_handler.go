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


// GetBooks godoc
// @Summary List all books
// @Description Get all books from database
// @Tags books
// @Produce json
// @Success 200 {array} domain.Book
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
    books, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
        return
    }
    c.JSON(http.StatusOK, books)
}


// GetBookByID godoc
// @Summary Get a book by ID
// @Description Get a single book by its UUID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} domain.Book
// @Failure 400 {object} map[string]string "Invalid book ID"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /book/{id} [get]
func (h *BookHandler) GetBookByID(c *gin.Context) {
    id := c.Param("id")

    book, err := h.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}


// CreateBook godoc
// @Summary Create a new book
// @Description Create a book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param book body domain.Book true "Book data"
// @Success 201 {object} domain.Book
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to create book"
// @Router /book [post]
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

// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update book identified by UUID with new data
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body domain.Book true "Updated book data"
// @Success 200 {object} domain.Book
// @Failure 400 {object} map[string]string "Invalid input or ID"
// @Failure 404 {object} map[string]string "Book not found or update failed"
// @Router /book/{id} [put]
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

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book by UUID
// @Tags books
// @Param id path string true "Book ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid book ID or delete failed"
// @Router /book/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
    id := c.Param("id")
    err := h.service.Delete(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID or failed to delete"})
        return
    }
    c.Status(http.StatusNoContent)
}
