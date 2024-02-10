package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/comments"

	resp, err := http.Get(url)
	if err != nil {
		//if request fails:
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//StatusOK = 200
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body) //return []byte, then cast to string
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	var jsonData map[string]interface{}

	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

}
