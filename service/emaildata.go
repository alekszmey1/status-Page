package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func NewEmailData(str []string) *EmailData {
	ed := EmailData{}
	ed.Country = str[0]
	ed.Provider = str[1]
	ed.DeliveryTime = helpers.StringToint(str[2])
	return &ed
}

func Email() []EmailData {

	providers := []string{"Orange", "Comcast", "AOL", "Gmail", "Yahoo", "Hotmail", "MSN", "Live", "RediffMail", "GMX",
		"Protonmail", "Yandex", "Mail.ru"}
	countriesString := helpers.CountryString()

	var storageED []EmailData
	smsDataCSV := "../StatusPage//simulator/email.data"
	smsDataString := helpers.CsvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 3)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 1)
	splitStrings = helpers.ExaminationCountry(splitStrings, countriesString)

	for _, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewEmailData(s)
		storageED = append(storageED, *l)
	}
	return storageED
	//fmt.Println(storageED)

}
