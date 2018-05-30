package main
import (
    "io"
    "net/http"
)

func FirstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "First Page!")
}

func SecondPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Second Page!")
}

func main() {
    http.HandleFunc("/", FirstPage)
    http.HandleFunc("/second", SecondPage)
    http.ListenAndServe(":8081", nil)
}
