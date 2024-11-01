package main

import (
    "fmt"
    "log"
    "net/http"
    "myapp/routes" // pastikan untuk mengganti dengan path yang sesuai
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "time"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    db := client.Database("myapp")

    r := routes.SetupRoutes(db)
    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
