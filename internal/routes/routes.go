package routes

import (
	"arthamna/rplLibrary/internal/handlers"
	"arthamna/rplLibrary/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController handlers.UserHandler, bookController handlers.BookHandler, categoryController handlers.CategoryHandler) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
		auth.POST("/admin/register", userController.Register)
	}

	// user
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/profile", userController.UploadPicture)

		// categories
		api.GET("/categories", categoryController.GetAllCategories)
	}

	books := r.Group("/books")
	books.Use(middleware.AuthMiddleware())
	{
		// borrow book
		books.POST("/borrow", bookController.BorrowBook)
		// filtering
		books.GET("/category/:category", bookController.FindByCategory)
		books.GET("/status/:status", bookController.FindByStatus)
	}

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// books
		// crud
		admin.POST("/book", bookController.CreateBook)
		admin.GET("/books", bookController.GetAllBooks)
		admin.GET("/book/:id", bookController.GetBook)
		admin.PUT("/book/:id", bookController.UpdateBook)
		admin.DELETE("/book/:id", bookController.DeleteBook)

		// uplod
		admin.POST("/books/profile", bookController.UploadBookPicture)
		// update book returned
		admin.POST("/books/returned/:id", bookController.SetReturnedBook)

		// Category
		// crud
		admin.POST("/category", categoryController.CreateCategory)
		admin.GET("/categories", categoryController.GetAllCategories)
		admin.GET("/category/:id", categoryController.GetCategory)
		admin.PUT("/category/:id", categoryController.UpdateCategory)
		admin.DELETE("/category/:id", categoryController.DeleteCategory)

	}
}
