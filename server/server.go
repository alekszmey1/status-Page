package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	mux := http.Server{
		Addr:    "127.0.0.1:8282",
		Handler: r,
	}
	r.HandleFunc("/", handleConnection)
	http.ListenAndServe(mux.Addr, r)

}
func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет"))

}
