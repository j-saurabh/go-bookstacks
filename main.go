package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func router() *mux.Router {

	rtr := mux.NewRouter()

	rtr.HandleFunc("/test", handler).Methods("GET")

	statFileDir := http.Dir("./assets/")

	statFileHandler := http.StripPrefix("/assets/",http.FileServer(statFileDir))

	rtr.PathPrefix("/assets/").Handler(statFileHandler).Methods("GET")


	rtr.HandleFunc("/book", getBooksHandler).Methods("GET")
	rtr.HandleFunc("/book", createBookHandler).Methods("POST")

	return rtr
}

func main() {

        rtr := router()

        err := http.ListenAndServe(":5000",rtr)

        if err != nil {

                panic(err.Error())
        }
}
