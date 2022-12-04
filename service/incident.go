package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"encoding/json"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"'`
}

func Incident() []IncidentData {
	url := "http://127.0.0.1:8383/accendent"
	//log.Println("открыли url " + url)
	incidentStorage, _ := createStorageIncident(url)
	//fmt.Println(incidentStorage)
	return incidentStorage
}

func createStorageIncident(url string) ([]IncidentData, error) {
	stringContent, err := helpers.UrlToString(url)
	stringContentSlice := helpers.StringToSliceString(stringContent)
	m := makeStorageIncident(stringContentSlice)
	cl := cleanSliceIncident(m)

	return cl, err
}
func makeStorageIncident(str []string) []*IncidentData {
	var SD []*IncidentData
	for _, s2 := range str {
		mms := createIncident([]byte(s2))
		SD = append(SD, mms)
	}
	//log.Println("заанмаршали каждое значение массива строк, создали срез структур формата mmsdata")
	return SD
}

func createIncident(b []byte) *IncidentData {
	var inc *IncidentData
	if err := json.Unmarshal(b, &inc); err != nil {
		//log.Printf("возникла ошибка в анмаршале %s ", err)
		inc = nil
	}
	return inc
}
func cleanSliceIncident(m []*IncidentData) []IncidentData {
	var n []IncidentData
	for _, val := range m {
		if val != nil {
			if val.Status == "active" || val.Status == "closed" {
				n = append(n, *val)
			}

		}
	}
	//log.Println("почистили слайс support от пустых срезов")
	return n
}
