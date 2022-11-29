package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"encoding/json"
	"fmt"
	"log"
)

type StorageData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"'`
}

func Support() {
	log.Println("создан  сервер")
	url := "http://127.0.0.1:8383/support"
	mmsStorage, _ := createStorageSupport(url)
	for _, data := range mmsStorage {
		fmt.Println(data)
	}
}

func createStorageSupport(url string) ([]*StorageData, error) {
	stringContent, err := helpers.UrlToString(url)
	stringContentSlice := helpers.StringToSliceString(stringContent)
	m := makeStorageSupport(stringContentSlice)
	cl := cleanSliceSupport(m)
	//fmt.Println(stringContent)
	return cl, err
}
func makeStorageSupport(str []string) []*StorageData {
	var SD []*StorageData
	for _, s2 := range str {
		mms := createSupport([]byte(s2))
		SD = append(SD, mms)
	}
	log.Println("заанмаршали каждое значение массива строк, создали срез структур формата mmsdata")
	return SD
}

func createSupport(b []byte) *StorageData {
	var sup *StorageData
	if err := json.Unmarshal(b, &sup); err != nil {
		log.Printf("возникла ошибка в анмаршале %s ", err)
		sup = nil
	}
	return sup
}
func cleanSliceSupport(m []*StorageData) []*StorageData {
	var n []*StorageData
	for _, val := range m {
		if val != nil {
			n = append(n, val)
		}
	}
	log.Println("почистили слайс support от пустых срезов")
	return n
}
