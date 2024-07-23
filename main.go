package main

import (
    "github.com/DraouiBilal/code-runner-backend-lib/routing"
    "github.com/DraouiBilal/code-runner-backend-lib/api"
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
       fmt.Println(req.PathValue(""))
        fmt.Fprint(w,"Testing if it works")
    }, []routing.Middleware{})
    server := routing.Server{}
    
    server.AddRouter(&dockerRouter)

    server.InitServer(&routing.Options{
		Middlewares: []routing.Middleware{routing.Middlewares.Logging},
	})

	log.Printf("Server starting on %s", server.FullAddr)

	//err := server.ListenAndServe()

	//if err != nil {
	//	fmt.Println(err)
	//}
    type Test struct {
        Test string `json:"test"`
    }
    response := api.MakeRequest[Test]("http://localhost:8080/docker/test/5","GET", struct{test string}{test: "test"},api.Options{})
    fmt.Println(response.Test)
}
