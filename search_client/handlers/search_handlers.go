package handlers

import (
    "fmt"
//	"strconv"
    "net/http"
    "context"
//    "encoding/json"

    "google.golang.org/grpc"
//    pb "../../api"
    reg "../../registry"
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

    top := r.URL.Query().Get("top")
    if top == "" {
        http.Error(w, "Expecting top in query", http.StatusBadRequest)
        return
    }

	//n, err := strconv.Atoi(top)
	//if err != nil {
    //    http.Error(w, err.Error(), http.StatusBadRequest)
    //    return
	//}

    //backend := "localhost:8181"

    conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
    if err != nil{
        http.Error(w, "Could not connect to service registry", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    srClient := reg.NewRegistryClient(conn)
    services,_ := srClient.GetAllServices(context.Background(), &reg.EmptyParam{})

    fmt.Fprint(w, fmt.Sprint("%v", services.Services))
    /*
    conn, err := grpc.Dial(backend, grpc.WithInsecure())
    if err != nil{
        http.Error(w, "Could not connect to search backend.", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    client := pb.NewDistanceClient(conn)

    query := &pb.Query{Id: id, Top: int32(n)}
    res, err := client.Dist(context.Background(), query)

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
	w.Write(jsonned)*/
}
