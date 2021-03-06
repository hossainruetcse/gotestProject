package server
import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Run() {
    http.HandleFunc("/hello", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}