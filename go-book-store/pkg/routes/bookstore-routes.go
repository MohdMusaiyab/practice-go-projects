package routes

import (
	"github.com/MohdMusaiyab/go-book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//For Creating a Book
	router.HandleFunc(("/book/"), controllers.CreateBook).Methods("POST")

	//For getting list of all the Books

	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")

	//For Getting a Single Book

	router.HandleFunc("/book/{bookid}", controllers.GetBook).Methods("GET")

	//For Deleting a Book

	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")

	//For Updating a Book

	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")

}
