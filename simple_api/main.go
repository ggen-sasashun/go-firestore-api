package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"example.com/simpleapi/handlers"
)

const (
	PROJECT_ID = "sasashun"
)

func createClient(ctx context.Context) *firestore.Client {

	client, err := firestore.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func main() {
	log.Println("starting server...")
	log.Println("project id is", PROJECT_ID)

	// ハンドラの登録
	http.HandleFunc("/hello", handlers.Hello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// サーバの起動
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
