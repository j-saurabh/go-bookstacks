package main

import(
	"net/http"
	"github.com/gorilla/mux"
)


func router() *mux.Router {

	rtr := mux.NewRouter()

	statFileDir := http.Dir("./books/")

	statFileHandler := http.StripPrefix("/books/",http.FileServer(statFileDir))

	rtr.PathPrefix("/books/").Handler(statFileHandler).Methods("GET")


	rtr.HandleFunc("/list", getBooksHandler).Methods("GET")
	rtr.HandleFunc("/list", createBookHandler).Methods("POST")

	return rtr
}

func main() {

        rtr := router()

        err := http.ListenAndServe(":5000",rtr)

        if err != nil {

                panic(err.Error())
        }
}
