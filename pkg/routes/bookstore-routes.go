package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sonochiwa/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r chi.Router) {
	r.Get("/book", controllers.GetAllBooks)
	r.Get("/book/{bookId}", controllers.GetBookById)
	r.Post("/book", controllers.CreateBook)
	r.Put("/book/{bookId}", controllers.UpdateBook)
	r.Delete("/book/{bookId}", controllers.DeleteBook)
}
