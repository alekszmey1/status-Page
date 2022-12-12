package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"'`
}

func Support() ([]SupportData, error) {
	log.Info("Получаем данные support")
	url := "http://127.0.0.1:8383/support"
	st, err := createStorageSupport(url)
	log.Info("Получены данные support")
	return st, err
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
	log.Info("заанмаршали каждое значение массива строк, создали срез структур формата support")
	return SD
}

func createSupport(b []byte) *SupportData {
	var sup *SupportData
	if err := json.Unmarshal(b, &sup); err != nil {
		log.Printf("возникла ошибка в анмаршале %s ", err)
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
	log.Info("почистили слайс support от пустых срезов")
	return n
}
