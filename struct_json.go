package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jsonBlob := []byte(`{
		"width": 48,
		"length": 64,
		"border": "solid"
}`)

	type rectangle struct {
		Width  int    `json:"width"`
		Length int    `json:"length"`
		Border string `json:"border"`
	}

	var rec rectangle

	err := json.Unmarshal(jsonBlob, &rec)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rec)
}

// END OMIT
