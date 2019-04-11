package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if len(r.URL.Path) > 1 {
        fmt.Fprintf(w, "path = %s\n", r.URL.Path[1:])
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
