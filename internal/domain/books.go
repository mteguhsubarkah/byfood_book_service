package domain

import (
    "github.com/google/uuid"
    "time"
	"gorm.io/gorm"

)

type Book struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Title     string    `gorm:"not null" json:"title"`
    Author    string    `gorm:"not null" json:"author"`
    Year      int       `gorm:"not null" json:"year"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

var input struct {
    Title  string `json:"title" binding:"required,min=1"`        // at least 1 character
    Author string `json:"author" binding:"required,min=1"`       // at least 1 character
    Year   int    `json:"year" binding:"required,gte=0,lte=2100"` // reasonable year range
}


// BeforeCreate hook, ensuring the ID is generated
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
    if b.ID == uuid.Nil {
        b.ID = uuid.New()
    }
    return
}
