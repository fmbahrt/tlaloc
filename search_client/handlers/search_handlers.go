package handlers

import (
    "fmt"
    "net/http"
    "context"
    "encoding/json"

    "google.golang.org/grpc"
    pb "../../api"
)

// Move this
func Index(w http.ResponseWriter, r *http.Request) {
    // Set headers and return type ALWAYS
    fmt.Fprintln(w, "WELCOME!")
}

func Search(w http.ResponseWriter, r *http.Request) {

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Expecting id in query", http.StatusBadRequest)
        return
    }

    backend := "localhost:8181"

    conn, err := grpc.Dial(backend, grpc.WithInsecure())
    if err != nil{
        http.Error(w, "Could not connect to search backend.", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    client := pb.NewDistanceClient(conn)

    queryId := &pb.UniqueId{Id: id}
    res, err := client.Dist(context.Background(), queryId)

    if err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jsonned, err := json.Marshal(res)
    if err != nil{
        http.Error(w, "Cannot marshal response", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonned)
}
