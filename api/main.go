package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type City struct {
	Name string `json:"city"`

	Rent float32 `json:"rent"`
	// insert data types that you wanna add

}

func get_database() []City {
	var body []City
	file, err := os.Open("database.json")
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(file).Decode(&body)
	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
	return body
}

func front(w http.ResponseWriter, r *http.Request) {
	var data []City = get_database()
	var res []byte
	city1, err := r.URL.Query()["city1"]
	if !err {
		log.Fatal(err)
	}
	city2, _ := r.URL.Query()["city2"]
	for i := 0; i < len(data); i++ {
		if city1[0] == data[i].Name || city2[0] == data[i].Name {
			res, _ = json.Marshal(data[i])
			w.Write(res)
		}
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		file, err := ioutil.ReadFile("database.json")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(file)
	}
}

func main() {
	http.HandleFunc("/front", front)
	http.HandleFunc("/admin", admin)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":6969", nil); err != nil {
		log.Fatal(err)
	}
}
