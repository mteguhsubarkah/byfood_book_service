package service

import (
    "github.com/google/uuid"
	"byfood_service/internal/domain"
    "gorm.io/gorm"
)

type BookService struct {
    db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
    return &BookService{db: db}
}

func (s *BookService) GetAll() ([]domain.Book, error) {
    var books []domain.Book
    err := s.db.Find(&books).Error
    return books, err
}

func (s *BookService) GetByID(id string) (*domain.Book, error) {
    uuidID, err := uuid.Parse(id)
    if err != nil {
        return nil, err
    }
    var book domain.Book
    err = s.db.First(&book, "id = ?", uuidID).Error
    if err != nil {
        return nil, err
    }
    return &book, nil
}

func (s *BookService) Create(book *domain.Book) error {
    return s.db.Create(book).Error
}

func (s *BookService) Update(id string, updatedData *domain.Book) (*domain.Book, error) {
    uuidID, err := uuid.Parse(id)
    if err != nil {
        return nil, err
    }
    var book domain.Book
    if err := s.db.First(&book, "id = ?", uuidID).Error; err != nil {
        return nil, err
    }

    book.Title = updatedData.Title
    book.Author = updatedData.Author
    book.Year = updatedData.Year

    if err := s.db.Save(&book).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func (s *BookService) Delete(id string) error {
    uuidID, err := uuid.Parse(id)
    if err != nil {
        return err
    }
    return s.db.Delete(&domain.Book{}, "id = ?", uuidID).Error
}
