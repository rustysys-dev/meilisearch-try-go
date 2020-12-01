package main

import (
	"fmt"
	"github.com/joho/godotenv"
	ms "github.com/meilisearch/meilisearch-go"
	"log"
	"os"
	. "rustysys-dev/m/v2/internal/ok"
	"rustysys-dev/m/v2/internal/pp"
)

func main() {
	log.Print("Retrieving deserts from index")

	err := godotenv.Load()
	Ok(err)

	var client = ms.NewClient(ms.Config{
		Host: os.Getenv("MEILISEARCH_HOST"),
	})

	result, err := client.Search("uid").Search(ms.SearchRequest{
		Query: "desert",
	})
	Ok(err)

	json_bytes, err := pp.MarshalJSON(result.Hits)
	Ok(err)

	fmt.Printf(string(json_bytes))
}
