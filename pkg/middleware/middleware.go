package middleware

import (
	"fmt"
	"net/http"
)

func Middleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Middleware ", r.URL)
		h.ServeHTTP(w, r)
	})
}
