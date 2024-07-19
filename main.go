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
    dockerRouter := routing.Router{
    }	

    dockerRouter.Get("/test/{id}", func (w http.ResponseWriter, req *http.Request){
        
        fmt.Println("Query")
        fmt.Println(req.URL.Parse(""))
        fmt.Fprint(w,"Testing if it works")
    }, []routing.Middleware{})

    server := routing.Server{}
    
    server.AddRouter(&dockerRouter)

    server.InitServer(&routing.Options{
		Middlewares: []routing.Middleware{loggingMiddleware},
	})

	log.Printf("Server starting on %s",server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
