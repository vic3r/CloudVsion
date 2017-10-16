package main

import (
	"os"
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	//"image"
	//"bytes"
	//"image/jpeg"
	//"encoding/base64"
	//"html/template"
)

type Image struct {
	  os.File
}

func passImage(w http.ResponseWriter, img *image.Image){
	print(w);
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	router := mux.NewRouter()
	router.HandleFunc("/Image/{id}", passImage).Methods("POST");



	log.Fatal(http.ListenAndServe(":8000", nil))

}
