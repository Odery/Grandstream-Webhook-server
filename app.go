// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

var telegramBot *TeleBot
var callCount map[string]int

func main() {
	var err error
	telegramBot, err = NewTeleBot()
	if err != nil {
		log.Fatal(err)
	}

	callCount = make(map[string]int)

	mux := http.NewServeMux()
	mux.HandleFunc("/siphook", sipHook)
	mux.HandleFunc("/", indexPageHandler)
	http.ListenAndServe("0.0.0.0:80", mux)
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func sipHook(w http.ResponseWriter, r *http.Request) {
	remote := r.URL.Query().Get("remote")
	callCount[remote] += 1
	msg := fmt.Sprintf("- Incoming call, number: %v, total calls from this number: %v", remote, callCount[remote])

	log.Println(msg)

	err := telegramBot.sendMsg(msg)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
	}

	w.WriteHeader(http.StatusOK)
}
