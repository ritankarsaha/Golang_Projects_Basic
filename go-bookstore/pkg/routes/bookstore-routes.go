package routes
import (
	"github.com/gorilla/mux"
	"github.com/ritankarsaha/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}",controllers.DeleteBook).Methods("DELETE")
}

