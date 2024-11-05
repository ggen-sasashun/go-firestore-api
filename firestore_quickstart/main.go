package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func createClient(ctx context.Context) *firestore.Client {

	// Google Cloud のプロジェクトID（デフォルト値）
	projectID := "sasashun"

	// プロジェクト ID を実行時に --project フラグで指定する
	flag.StringVar(&projectID, "project", projectID, "The Google Cloud Platform project ID.")
	flag.Parse()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// client が使い終わったら必ず Close する
	// defer client.Close()
	return client
}

func main() {

	// Firestore クライアントを作成
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	// Firestore にデータを追加1
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	// Firestore にデータを追加2
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}

	// Firestore からデータを読み取り
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}
