package dtos

import (
	"mime/multipart"
	"time"
)

type BookCreateRequest struct {
	Author      string   `json:"author" binding:"required"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	CategoryIDs []string `json:"category_ids" binding:"required"`
}

type UploadBookPictureRequest struct {
	BookID      string                `form:"book_id" binding:"required"`
	BookPicture *multipart.FileHeader `form:"book_picture" binding:"required"`
}

type BookUpdateRequest struct {
	BookID      string   `json:"book_id" binding:"required"`
	Author      string   `json:"author"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CategoryIDs []string `json:"category_ids"`
}

type BorrowBookRequest struct {
	BookID string `json:"book_id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}


//

type BorrowBookResponse struct {
	BookID     string    `json:"book_id"`
	Title      string    `json:"title"`
	UserID     string    `json:"user_id"`
	Username   string    `json:"username"`
	BorrowedAt time.Time `json:"borrowed_at"`
}

type SetBookReturnedResponse struct {
	BookID     string    `json:"book_id"`
	BorrowingID     string    `json:"borrowing_id"`
	ReturnedAt time.Time `json:"returned_at"`
}

type BookResponse struct {
	BookID      string    `json:"book_id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BookPicture string    `json:"book_picture,omitempty"`
	Status      string    `json:"status"`
	CategoryIDs []string  `json:"category_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UploadBookPictureResponse struct {
	BookID      string `json:"book_id"`
	BookPicture string `json:"book_picture"`
}
