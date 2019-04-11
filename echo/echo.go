package main

import (
    "fmt"
    "log"
    "net/http"
    "io"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println(r)
    if len(r.Header.Get("X-Echo")) > 1 {
        fmt.Fprintf(w, "header = %s\n", r.Header.Get("X-Echo"))
    }
    if len(r.URL.Path) > 1 {
        fmt.Fprintf(w, "path = %s\n", r.URL.Path[1:])
    }
    if len(r.URL.RawQuery) > 1 {
        fmt.Fprintf(w, "query = %s\n", r.URL.RawQuery)
    }
    if len(r.URL.Fragment) > 1 {
        fmt.Fprintf(w, "fragment = %s\n", r.URL.Fragment)
    }
    hasBody := []string{"POST","PUT","PATCH"}
    if in(r.Method, hasBody) {
        io.Copy(w, r.Body)
    }
}

func in(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
