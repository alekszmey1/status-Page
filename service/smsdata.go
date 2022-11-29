package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func NewSMSData(str []string) *SMSData {
	sms := SMSData{}
	sms.Country = str[0]
	sms.Bandwidth = str[1]
	sms.ResponseTime = str[2]
	sms.Provider = str[3]
	return &sms
}

type StorageSD struct {
	storageSMSData map[int]*SMSData
}

func SmsData() {
	var storageSMS []SMSData
	providers := []string{"Topol", "Rond", "Kildy"}
	countriesString := helpers.CountryString()
	smsDataCSV := "../StatusPage/simulator/sms.data"
	smsDataString := helpers.CsvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 4)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 3)
	splitStrings = helpers.ExaminationCountry(splitStrings, countriesString)
	for _, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewSMSData(s)
		storageSMS = append(storageSMS, *l)
	}

}
