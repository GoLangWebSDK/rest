package main

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running Logger Middleware for Request: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Running Authorization...")

		secureRoutes := []string{
			"/api/posts/admin",
		}

		for _, route := range secureRoutes {
			if r.URL.Path == route {
				fmt.Println("User is authorized.")
				next.ServeHTTP(w, r)
				return
			}
		}

		fmt.Println("Path is public.")
		next.ServeHTTP(w, r)
	})
}
