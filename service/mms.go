package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type storageMMS struct {
	storage []*MMSData
}

func (u *storageMMS) put(mms *MMSData /*i int*/) {
	u.storage = append(u.storage, mms)
}

func NewStorageMMSD() *storageMMS {
	return &storageMMS{storage: []*MMSData{}}
}
func (s *storageMMS) Get(w http.ResponseWriter, r *http.Request) {
	providers := []string{"Topol", "Rond", "Kildy"}
	countriesString := helpers.CountryString()
	var mms MMSData
	if r.Method == "GET" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		if err := json.Unmarshal(content, &mms); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		checkCountry := helpers.CheckCoutry(countriesString, mms.Country)
		checkProviders := helpers.CheckProviders(mms.Provider, providers)
		if checkCountry == true && checkProviders == true {
			s.put(&mms)
		}
		w.WriteHeader(http.StatusCreated)
	}
	//w.WriteHeader(http.StatusBadRequest)
}

func (s *storageMMS) GetAll(w http.ResponseWriter, r *http.Request) { // возвращает всех пользователей
	if r.Method == "POST" {
		//response := ""
		for _, user := range s.storage { //итерируемся по всем пользователям в mape
			fmt.Println(user)
		}
		w.WriteHeader(http.StatusOK)
		//w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func MmsData() {
	MMSD := NewStorageMMSD()
	mux := http.NewServeMux()
	mux.HandleFunc("/get", MMSD.Get)
	mux.HandleFunc("/getall", MMSD.GetAll)
	http.ListenAndServe("localhost:8383", mux)
	fmt.Println(MMSD)

}
