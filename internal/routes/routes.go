package routes

import (
	"arthamna/rplLibrary/internal/handlers"
	"arthamna/rplLibrary/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController handlers.UserHandler, bookController handlers.BookHandler, categoryController handlers.CategoryHandler) {
	
	// auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
		auth.POST("/admin/register", userController.RegisterAdmin)
	}

	// user
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/profile", userController.UploadPicture)
		
		// kinda confused, but let's say user dan admin punya endpoint tersendiri untuk bisa mengakses categories 
		// sehingga dapat melakukan filtering dengan lebih mudah
		books := api.Group("/books")
		{
			// borrow book
			books.POST("/borrow", bookController.BorrowBook)
			books.POST("/borrows", bookController.BorrowMultipleBook)

			// filtering
			books.GET("/status/:status", bookController.FindByStatus)
			// categories, untuk mencari buku berdasarkan kategori
			books.GET("/categories", categoryController.GetAllCategories)
			books.GET("/category/:category", bookController.FindByCategory)
			// searching : /books/search?q=harry
			books.GET("/search", bookController.SearchBooks)
		}

	}

	// Admin 
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
		admin.POST("/book/cover", bookController.UploadBookPicture)
		// update book returned
		admin.POST("/book/returned", bookController.SetReturnedBook)
		admin.POST("/books/returned", bookController.SetMultipleReturnedBook)

		// Category
		// crud
		admin.POST("/category", categoryController.CreateCategory)
		admin.GET("/categories", categoryController.GetAllCategories)
		admin.GET("/category/:id", categoryController.GetCategory)
		admin.PUT("/category/:id", categoryController.UpdateCategory)
		admin.DELETE("/category/:id", categoryController.DeleteCategory)

	}
}
