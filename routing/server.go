package routing

import (
    "net/http"
    "strconv"
    "time"
    "log"
)

type Options struct {
    Host string
    Port int
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}

func createHandler() (http.Handler) {
    handler := http.NewServeMux()
    loggedHandler := loggingMiddleware(handler)
    return loggedHandler
}

func CreateServer (options *Options) (*http.Server) {
    handler := createHandler()
    
    address := options.Host + ":" + strconv.Itoa(options.Port)

    server := http.Server {
        Addr: address,
        Handler: handler,
    }
    
    return &server
}

