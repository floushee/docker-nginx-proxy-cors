package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var tom *person = &person{
	Name: "floushee",
	Age:  32,
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(tom)
	w.Write(j)
}

func main() {
	http.HandleFunc("/", apiHandler)

	log.Println("Go!")
	http.ListenAndServe(":3000", nil)
}
