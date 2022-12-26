package server

import (
	"awesomeProject/skillbox/StatusPage/service"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func App() {
	r := mux.NewRouter()

	mux := http.Server{
		Addr:    "127.0.0.1:8282",
		Handler: r,
	}
	r.Handle("/", http.FileServer(http.Dir("./web")))
	r.HandleFunc("/api", handleConnection)
	http.ListenAndServe(mux.Addr, r)

}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	m := service.MakeResultT()
	data, err := json.Marshal(m)
	if err != nil {
		log.Info(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
