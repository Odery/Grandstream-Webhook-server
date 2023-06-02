package main

import (
	"net/http"
	"log"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/siphook", sipHook)
	mux.HandleFunc("/", indexPageHandler)

	http.ListenAndServe("0.0.0.0:80", mux)
}

func indexPageHandler(w http.ResponseWriter , r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func sipHook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	remote := r.URL.Query().Get("remote")

	log.Println("- Incoming call, number: ", remote)
}