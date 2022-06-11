package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func get_database() ([]City, error) {
	var body []City
	file, err := os.Open("database.json")
	if err != nil {
		return nil, err
	}
	json.NewDecoder(file).Decode(&body)
	if err = file.Close(); err != nil {
		return nil, err
	}
	return body, err
}

func modify_database(city City) error {
	data, err := get_database()
	if err != nil {
		return err
	}
	for i := 0; i < len(data); i++ {
		if city.Name == data[i].Name {
			data[i] = city
		}
	}
	file, err := os.Create("database.json")
	if err != nil {
		return err
	}
	defer file.Close()
	b, _ := json.Marshal(data)
	file.Write(b)
	return nil
}

func front(w http.ResponseWriter, r *http.Request) {
	data, err := get_database()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var res []byte
	city1, found := r.URL.Query()["city1"]
	if !found {
		w.WriteHeader(400)
		return
	}
	city2, _ := r.URL.Query()["city2"]
	var final []City
	for i := 0; i < len(data); i++ {
		if city1[0] == data[i].Name || city2[0] == data[i].Name {
			final = append(final, data[i])
		}
	}
	res, _ = json.Marshal(final)
	w.Write(res)
}

func admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		file, err := ioutil.ReadFile("database.json")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(file)
	}
	if r.Method == "POST" {
		var change City
		json.NewDecoder(r.Body).Decode(&change)
		err := modify_database(change)
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}
}

func main() {
	http.HandleFunc("/front", front)
	http.HandleFunc("/admin", admin)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":6969", nil); err != nil {
		os.Exit(500)
	}
}
