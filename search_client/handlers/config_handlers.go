package handlers

import (
    "fmt"
    "net/http"
)

func Config(w http.ResponseWriter, r *http.Request) {
    // Set headers and return type ALWAYS
    fmt.Fprintln(w, "Application Config")
}
