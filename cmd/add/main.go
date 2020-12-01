package main

import (
	"encoding/json"
	"time"

	"github.com/joho/godotenv"
	ms "github.com/meilisearch/meilisearch-go"
	"io/ioutil"
	"log"
	"os"
	. "rustysys-dev/m/v2/internal/ok"
)

func main() {
	log.Print("Adding entries to index")

	err := godotenv.Load()
	Ok(err)

	var client = ms.NewClient(ms.Config{
		Host: os.Getenv("MEILISEARCH_HOST"),
	})

	jsonFile, err := os.Open("data.json/docs.json")
	Ok(err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	Ok(err)
	var docs []map[string]interface{}
	err = json.Unmarshal(byteValue, &docs)
	Ok(err)

	updateDocsResponse, err := client.Documents("uid").AddOrUpdate(docs)
	Ok(err)

	time.Sleep(1 * time.Second)

	retrieveData, err := client.Updates("uid").Get(updateDocsResponse.UpdateID)
	Ok(err)

	log.Print("UpdateID: ", updateDocsResponse.UpdateID)

	log.Print(retrieveData)
}
