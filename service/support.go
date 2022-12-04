package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"encoding/json"
	"log"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"'`
}

func Support() []SupportData {
	log.Println("создан  сервер")
	url := "http://127.0.0.1:8383/support"
	st, _ := createStorageSupport(url)
	return st
}

func createStorageSupport(url string) ([]SupportData, error) {
	stringContent, err := helpers.UrlToString(url)
	stringContentSlice := helpers.StringToSliceString(stringContent)
	m := makeStorageSupport(stringContentSlice)
	cl := cleanSliceSupport(m)
	return cl, err
}
func makeStorageSupport(str []string) []*SupportData {
	var SD []*SupportData
	for _, s2 := range str {
		mms := createSupport([]byte(s2))
		SD = append(SD, mms)
	}
	//log.Println("заанмаршали каждое значение массива строк, создали срез структур формата mmsdata")
	return SD
}

func createSupport(b []byte) *SupportData {
	var sup *SupportData
	if err := json.Unmarshal(b, &sup); err != nil {
		//log.Printf("возникла ошибка в анмаршале %s ", err)
		sup = nil
	}
	return sup
}
func cleanSliceSupport(m []*SupportData) []SupportData {
	var n []SupportData
	for _, val := range m {
		if val != nil {
			n = append(n, *val)
		}
	}
	//log.Println("почистили слайс support от пустых срезов")
	return n
}
