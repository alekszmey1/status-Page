/*package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type storageSupport struct {
	storage []*SupportData
}

func (u *storageSupport) put(sup *SupportData /*i int) {
	u.storage = append(u.storage, sup)
}

func NewStorageSD() *storageSupport {
	return &storageSupport{storage: []*SupportData{}}
}
func (s *storageSupport) Get(w http.ResponseWriter, r *http.Request) {
	//providers := []string{"Topol", "Rond", "Kildy"}
	//countriesString := helpers.CountryString()
	var sup SupportData
	if r.Method == "GET" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		if err := json.Unmarshal(content, &sup); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		//checkCountry := helpers.CheckCoutry(countriesString, mms.Country)
		//checkProviders := helpers.CheckProviders(mms.Provider, providers)
		//if checkCountry == true && checkProviders == true {
		s.put(&sup)
		//}
		w.WriteHeader(http.StatusCreated)
	}
	//w.WriteHeader(http.StatusBadRequest)
}

func (s *storageSupport) GetAll(w http.ResponseWriter, r *http.Request) { // возвращает всех пользователей
	if r.Method == "POST" {
		for _, user := range s.storage { //итерируемся по всем пользователям в mape
			fmt.Println(user)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func mms() {
	SD := NewStorageSD()
	mux := http.NewServeMux()
	mux.HandleFunc("/get", SD.Get)
	mux.HandleFunc("/getall", SD.GetAll)
	http.ListenAndServe("localhost:8484", mux)
	fmt.Println(SD)

}
*/