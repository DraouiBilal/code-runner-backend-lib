package main

import (
	"code-runner/backend/lib/routing"
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {
	host := "localhost"
	port := 8080
	server := routing.CreateServer(&routing.Options{
		Host:        host,
		Port:        port,
		Middlewares: []routing.Middleware{loggingMiddleware},
	})

	log.Printf("Server starting on %s:%v",host,port)
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
