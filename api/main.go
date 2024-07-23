package api

import (
    "bytes"
    "net/http"
    "encoding/json"
    "log" 
)

type Options struct {
    headers  map[string] string
}

func createRequest(url string, method string, body interface{}, options Options) *http.Request {
    payload, err := json.Marshal(body) 

    if err != nil {
        log.Fatal(err)
    }

    req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
    req.Header.Set("Content-Type", "applciation/json")
    
    for _,k := range options.headers {
        req.Header.Set(k,options.headers[k])
    }
    
    if err != nil {
        log.Fatal(err)
    }

    return req
}

func MakeRequest[T any](url string, method string, body interface{}, options Options) *T {
    client := &http.Client{}

    req := createRequest(url, method, body, options)

    res, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()

    if err != nil {
        log.Fatal(err)
    }

    responseJson := new(T)

    decoder := json.NewDecoder(res.Body)
    
    err = decoder.Decode(&responseJson)

    if err != nil {
        log.Fatal(err)
    }
    return responseJson
}
