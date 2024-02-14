package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Comment struct {
	postId int
	id     int
	name   string
	email  string
	body   string
}

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

	//Parse JSON data. Json to Struct unmarshal | Struct to json marshal

	var comment []Comment

	err = json.Unmarshal(data, &comment)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	//Create JSON file
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optionally set indentation for better readability
	if err := encoder.Encode(comment); err != nil {
		fmt.Printf("Error encoding JSON to file: %v\n", err)
		return
	}

	fmt.Println("JSON data fetched and saved to data.json")

}
