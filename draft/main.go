package main

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	if r.Method == "GET" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		stringContent := string(content)
		s3 := strings.Trim(stringContent, "{}")
		stringSliceContent := strings.Split(s3, ",")
		i := 0
		var miniStringSlice []string
		for _, s2 := range stringSliceContent {
			miniStringSlice = append(miniStringSlice, s2)
			if i == 3 {
				mmsString := "{" + strings.Join(miniStringSlice, ",") + "}"
				mms := createMMS([]byte(mmsString))
				checkCountry := helpers.CheckCoutry(countriesString, mms.Country)
				checkProviders := helpers.CheckProviders(mms.Provider, providers)
				if checkCountry == true && checkProviders == true {
					s.put(mms)
				}
				miniStringSlice = nil
				i = 0
				continue
			}
			i++
		}
		defer r.Body.Close()
		w.WriteHeader(http.StatusCreated)
	}
	fmt.Println(s)
}

func (s *storageMMS) GetAll(w http.ResponseWriter, r *http.Request) { // возвращает всех пользователей
	if r.Method == "POST" {
		for _, user := range s.storage { //итерируемся по всем пользователям в mape
			fmt.Println(user)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	MMSD := NewStorageMMSD()
	mux := http.NewServeMux()
	mux.HandleFunc("/get", MMSD.Get)
	//mux.HandleFunc("/getall", MMSD.GetAll)
	http.ListenAndServe("localhost:8080", mux)
}

func createMMS(b []byte) *MMSData {
	var mms *MMSData
	if err := json.Unmarshal(b, &mms); err != nil {
		mms = nil
	}
	return mms
}
