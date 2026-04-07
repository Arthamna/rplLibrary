package models

import "time"

type BookCategory struct {
	BookCategoryID string  `gorm:"column:book_category_id;primaryKey"`
	BookID         string  `gorm:"column:book_id;index"`
	CategoryID     string  `gorm:"column:category_id;index"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}