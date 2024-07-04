package main

import (
	"encoding/gob"
	"encoding/json"
  "fmt"
	"os"
)

func main() {
	type CharInfo struct {
		Char string   `json:"char"`
		Key  []string `json:"key"`
		// Name []string `json:"name"`
	}

	type Dataset struct {
		A []CharInfo `json:"A"`
		B []CharInfo `json:"B"`
		C []CharInfo `json:"C"`
		N []CharInfo `json:"N"`
	}

	jsonData, err := os.ReadFile("./data/cangjie.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var data Dataset
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	saveToGobFile("./data/data.gob", data)

	var loaded_data Dataset
	loadFromGobFile("./data/data.gob", &loaded_data)

	fmt.Println("Loaded_data: ", loaded_data)
}

func saveToGobFile(filename string, data interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}

func loadFromGobFile(filename string, data interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		panic(err)
	}
}
