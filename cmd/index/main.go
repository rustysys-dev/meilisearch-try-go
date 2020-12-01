package main

import (
	"github.com/joho/godotenv"
	ms "github.com/meilisearch/meilisearch-go"
	"log"
	"os"
	. "rustysys-dev/m/v2/internal/ok"
)

func main() {
	Ok(godotenv.Load())
	var client = ms.NewClient(ms.Config{
		Host: os.Getenv("MEILISEARCH_HOST"),
	})

	createIndexResponse, err := client.Indexes().Create(ms.CreateIndexRequest{
		Name:       "Primary Index",
		UID:        "uid",
		PrimaryKey: "id",
	})
	Ok(err)

	log.Print(createIndexResponse)

	indexList, err := client.Indexes().List()
	Ok(err)

	for _, v := range indexList {
		log.Print("Index: ", v)
	}
}
