package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	m "github.com/amrali/golang-blog/pkg/middleware"
	. "github.com/amrali/golang-blog/pkg/handler"
)


func HttpRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/register", m.Middleware(Register)).Methods("POST")
	r.HandleFunc("/login", m.Middleware(Login)).Methods("POST")
	r.HandleFunc("/logout", m.Middleware(Logout)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", r))
}