package routes

import (
	"github.com/gorilla/mux"
	"github.com/jackson-/golang-book-management/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook)
}
