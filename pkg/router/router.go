package router

import (
	"log"
	"net/http"

	. "github.com/amrali/golang-blog/pkg/handler"
	m "github.com/amrali/golang-blog/pkg/middleware"
	"github.com/gorilla/mux"
)

func HttpRequests() {
	r := mux.NewRouter().StrictSlash(true)

	// User-Authentication Section
	r.HandleFunc("/register", m.Middleware(Register)).Methods("POST")
	r.HandleFunc("/login", m.Middleware(Login)).Methods("POST")
	r.HandleFunc("/logout", m.Middleware(Logout)).Methods("POST")

	// Post Section
	r.HandleFunc("/api/newpost", m.Middleware(AddPost)).Methods("POST")
	r.HandleFunc("/api/post", m.Middleware(GetPost)).Methods("GET")
	r.HandleFunc("/api/updatepost", m.Middleware(UpdatePost)).Methods("PUT")
	r.HandleFunc("/api/deletepost", m.Middleware(DeletePost)).Methods("DELETE")

	// Comment-Section
	r.HandleFunc("/api/newcomment", m.Middleware(AddComment)).Methods("POST")
	r.HandleFunc("/api/comment", m.Middleware(GetComment)).Methods("GET")
	r.HandleFunc("/api/updatecomment", m.Middleware(UpdateComment)).Methods("PUT")
	r.HandleFunc("/api/deletecomment", m.Middleware(DeleteComment)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}
