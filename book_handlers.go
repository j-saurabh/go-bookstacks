package main

import(

	"fmt"
	"encoding/json"
	"net/http"

)

type Book struct {

	Name string `json:"bname"`
	Author string `json:"author"`
}

var books []Book

func getBooksHandler(w http.ResponseWriter, r *http.Request) {

	bookListBytes, err := json.Marshal(books)

	if err != nil {

		fmt.Println(fmt.Errorf("Error : %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bookListBytes)

}

func createBookHandler(w http.ResponseWriter, r *http.Request) {

	book := Book{}

	err := r.ParseForm()

	if err != nil {

		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	book.Name = r.Form.Get("bname")
	book.Author = r.Form.Get("author")


	books = append(books, book)

	http.Redirect(w, r, "/assets/", http.StatusFound)

}

