package routing

import (
    "net/http"
    "strconv"
)

type Middleware func(http.Handler) http.Handler

type Options struct {
    Host string
    Port int
    Middlewares []Middleware
}


func createHandler(middlewares []Middleware) (http.Handler) {
    mux := http.NewServeMux()
    var handler http.Handler = mux
    for _,middleware := range middlewares {
        handler = middleware(handler)
    }
    return handler
}

func CreateServer (options *Options) (*http.Server) {
    handler := createHandler(options.Middlewares)
    
    address := options.Host + ":" + strconv.Itoa(options.Port)

    server := http.Server {
        Addr: address,
        Handler: handler,
    }
    
    return &server
}

