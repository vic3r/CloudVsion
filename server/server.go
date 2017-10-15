package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	//"github.com/gorilla/mux"

)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	//router := mux.newRouter().StrictSlashtrue(True)
	//router.HandleFunc("/Image/{id}", CreateImage).Methods("POST");



	log.Fatal(http.ListenAndServe(":8000", nil))

}
