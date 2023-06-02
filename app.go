package main

import (
	"net/http"
	"log"
	"bufio"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/siphook", sipHook)

	http.ListenAndServe("80", mux)
}

func sipHook(w http.ResponseWriter, r *http.Request) {
	var bodyString []byte

	_, err := r.Body.Read([]byte(bodyString))
	if err != nil {
		log.Printf("Error reading body: %v", err)
	}

	log.Println(string(bodyString))
}