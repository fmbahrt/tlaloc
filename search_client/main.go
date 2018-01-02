package main

import (
    "log"
    "net/http"
    "./routing"
)

func main() {

    router := routing.NewRouter(routing.ServerRoutes)

    log.Fatal(http.ListenAndServe(":8080", router))
}
