package common

import (
	"encoding/json"
	"log"
)

func Println(v interface{}) {
	prettyJSON, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	log.Printf("%s\n", string(prettyJSON))
}
