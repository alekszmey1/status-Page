package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func MmsData() {
	log.Println("создан  сервер")
	url := "http://127.0.0.1:8383/mms"
	mmsStorage, _ := createStorageMMS(url)
	for _, data := range mmsStorage {
		fmt.Println(data)
	}
}

func createStorageMMS(url string) ([]*MMSData, error) {
	stringContent, err := helpers.UrlToString(url)
	stringContentSlice := helpers.StringToSliceString(stringContent)
	m := makeStorageMMS(stringContentSlice)
	c := cleanSliceMMS(m)
	return c, err
}
func makeStorageMMS(str []string) []*MMSData {
	var MD []*MMSData
	for _, s2 := range str {
		mms := createMMS([]byte(s2))
		MD = append(MD, mms)
	}
	log.Println("заанмаршали каждое значение массива строк, создали срез структур формата mmsdata")
	return MD
}

func stringToSliceString(s string) []string {
	s2 := strings.Trim(s, "[][]")
	s2 = strings.Replace(s2, "[", "", -1)
	s = strings.Replace(s2, "},{", "};{", -1)
	str := strings.Split(s, ";")
	log.Println("убрали лишние скобки, разбили строку на массив строк")
	return str
}

func createMMS(b []byte) *MMSData {
	var mms *MMSData
	if err := json.Unmarshal(b, &mms); err != nil {
		log.Printf("возникла ошибка в анмаршале %s ", err)
		mms = nil
	}
	return mms
}

func cleanSliceMMS(m []*MMSData) []*MMSData {
	countryString := helpers.CountryString()
	providers := []string{"Rond", "Topolo", "Kildy"}
	var n []*MMSData
	for _, val := range m {
		if val != nil {
			country := helpers.CheckCountry(countryString, val.Country)
			checkProviders := helpers.CheckProviders(val.Provider, providers)
			if country == true && checkProviders == true {
				n = append(n, val)
			}
		}
	}
	log.Println("почистили слайс mmsdata от пустых срезов, и проверили на соответствие странам и провайдерам")
	return n
}
