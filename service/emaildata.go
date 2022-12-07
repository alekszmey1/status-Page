package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"strings"

	log "github.com/sirupsen/logrus"
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
	log.Info("Получаем данные email")

	providers := []string{"Orange", "Comcast", "AOL", "Gmail", "Yahoo", "Hotmail", "MSN", "Live", "RediffMail", "GMX",
		"Protonmail", "Yandex", "Mail.ru"}
	countriesString := helpers.CountryString()

	var storageED []EmailData
	emailDataCSV := "./simulator/email.data"
	emailDataString := helpers.CsvInString(emailDataCSV)
	splitStrings := strings.Split(emailDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 3)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 1)
	splitStrings = helpers.ExaminationCountry(splitStrings, countriesString)

	for _, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewEmailData(s)
		storageED = append(storageED, *l)
	}
	log.Info("Получены данные email")
	return storageED

}
