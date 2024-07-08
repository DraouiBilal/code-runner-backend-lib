package main

import (
	"code-runner/backend/lib/routing"
	"fmt"
	"log"
	"net/http"
)


func main() {
    dockerRouter := routing.Router{
        Name: "docker",
    }	

   dockerRouter.Get("/test/{id}", func (w http.ResponseWriter, req *http.Request){
       fmt.Println(req.PathValue("id"))
       fmt.Println(req.URL.Query()["a"][0])
        fmt.Fprint(w,"Testing if it works")
    }, []routing.Middleware{})
    server := routing.Server{}
    
    server.AddRouter(&dockerRouter)

    server.InitServer(&routing.Options{
		Middlewares: []routing.Middleware{routing.Middlewares.Logging},
	})

	log.Printf("Server starting on %s", server.FullAddr)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
