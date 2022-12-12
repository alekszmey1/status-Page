package server

import (
	"awesomeProject/skillbox/StatusPage/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func App() {
	r := mux.NewRouter()
	mux := http.Server{
		Addr:    "127.0.0.1:8282",
		Handler: r,
	}
	r.HandleFunc("/api", handleConnection)
	http.ListenAndServe(mux.Addr, r)

}
func handleConnection(w http.ResponseWriter, r *http.Request) {
	m := service.MakeResultT()
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
