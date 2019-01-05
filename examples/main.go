package main

//go:generate go run generator.go

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("HTTP Server is running on port 8080")
	http.ListenAndServe(":8080", router())
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get", getHandler).Methods("GET")
	r.HandleFunc("/post", postHandler).Methods("POST")
	r.HandleFunc("/delete", deleteHandler).Methods("DELETE")
	return r
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You GET it!"))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You can POST it!"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You can DELETE it!"))
}
