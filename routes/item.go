package routes

import (
    "fmt"
    "net/http"
)

func ItemHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
