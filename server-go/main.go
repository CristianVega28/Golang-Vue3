package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"message"`
	Post string `json: "method"`
}

type Informacion struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi there")

	w.Header().Set("Content-Type", "application/json")

	hello := Message{Text: "Hello world", Post: r.Method}
	messagge, _ := json.Marshal(hello)

	fmt.Println(r.Method, http.MethodPost)
	if r.Method == http.MethodPost {
		var bodydata Informacion
		err := json.NewDecoder(r.Body).Decode(&bodydata)
		fmt.Println(bodydata.Name, bodydata.Job)
		if err != nil {
		}

	}
	json_string := string(messagge)
	w.Write([]byte(json_string))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
