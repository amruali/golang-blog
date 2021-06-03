package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func HttpRequests() {
	r := mux.NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe(":8080", r))
}