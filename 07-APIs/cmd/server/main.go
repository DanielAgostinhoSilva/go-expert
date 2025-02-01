package main

import (
	"github.com/DanielAgostinhoSilva/go-expert/07-APIs/config"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		config.Initialize()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(config.CFG.Database.DBDriver))
	})
	http.ListenAndServe(":8080", nil)
}
