package main

import (
	"fmt"
	"os"
	"encoding/json"
	"encoding/gob"
)

func main() {
	type char_data struct {
		Char string     `json:"char"`
		Key  []string `json:"key"`
		Name []string `json:"name"`
	}

	type dataset struct {
		A []char_data `json:"A"`
		B []char_data `json:"B"`
		C []char_data `json:"C"`
		N []char_data `json:"N"`
	}

	jsonData, err := os.ReadFile("./data/cangjie.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var data dataset
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

  saveToGobFile("./data/data.gob", data)

  var loaded_data dataset
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

