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
	log.Print("Retrieving entries from index")

	err := godotenv.Load()
	Ok(err)

	var client = ms.NewClient(ms.Config{
		Host: os.Getenv("MEILISEARCH_HOST"),
	})

	var a []interface{}
	Ok(client.Documents("uid").List(ms.ListDocumentsRequest{}, &a))

	json_bytes, err := pp.MarshalJSON(a)
	Ok(err)

	fmt.Printf(string(json_bytes))
}
