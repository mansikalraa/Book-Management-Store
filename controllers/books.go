package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mansikalra23/Project/BookManagement/models"
)

// Find all books
func FindBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	models.DB.Find(&books)

	json.NewEncoder(w).Encode(&books)
}

func FindBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.Book
	if err := models.DB.Where("id = ?", params["id"]).First(&book).Error; err != nil {
		fmt.Fprintf(w, "Record not found.")
		return
	}

	models.DB.First(&book, params["id"])
	json.NewEncoder(w).Encode(&book)
}

// Create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Validate input
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	models.DB.Create(&book)
	json.NewEncoder(w).Encode(&book)

	fmt.Fprintf(w, "Record created.")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", params["id"]).First(&book).Error; err != nil {
		fmt.Fprintf(w, "Record not found.")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	book = models.Book{ID: params["id"], Title: keyVal["title"], Author: keyVal["author"]}

	models.DB.Model(&book).Where("id = ?", params["id"]).Update(book)

	fmt.Fprintf(w, "Record updated.")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.Book
	if err := models.DB.Where("id = ?", params["id"]).First(&book).Error; err != nil {
		fmt.Fprintf(w, "Record not found.")
		return
	}

	models.DB.Where("id = ?", params["id"]).Delete(&book)

	fmt.Fprintf(w, "Record deleted.")
}
