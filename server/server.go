package server

import (
	"awesomeProject/skillbox/StatusPage/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
	//"github.com/gorilla/mux"
)

func App() {

	router := mux.NewRouter()
	router.HandleFunc("/api", handleConnection).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	http.Handle("/", router)
	http.ListenAndServe("127.0.0.1:8282", nil)
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
