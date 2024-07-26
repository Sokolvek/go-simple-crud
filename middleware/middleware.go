package middleware

import (
	"fmt"
	"net/http"
)

func UpdateDB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("middleware working!!!")
		next.ServeHTTP(w, r)
	})
}
func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		fmt.Println("middleware")
		next.ServeHTTP(w, r)
	})
}
