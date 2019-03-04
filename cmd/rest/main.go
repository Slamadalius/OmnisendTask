package main

import (
	"fmt"
	"log"
	"github.com/julienschmidt/httprouter"

	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Working")
}

func Filter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Working")
}

func main() {
	r := httprouter.New()
	r.GET("/", Index)

	fmt.Println("Server is listening on port :8080")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}