// BOOKS STORAGE CENTER USING GORILLA MUX AND MYSQL

package main

import (
	"net/http"
	"github.com/mansikalra23/Project/BookManagement/controllers"
	"github.com/mansikalra23/Project/BookManagement/models"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	models.ConnectDatabase()

	router := mux.NewRouter()
	router.HandleFunc("/books", controllers.FindBooks).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.FindBook).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
