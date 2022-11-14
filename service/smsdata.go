package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"fmt"
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

func NewStorageSD() *StorageSD {
	return &StorageSD{storageSMSData: make(map[int]*SMSData)}
}
func (u *StorageSD) put(sms *SMSData, i int) {
	u.storageSMSData[i] = sms
}
func (u *StorageSD) getAll() []*SMSData {
	var smsDats []*SMSData
	for _, v := range u.storageSMSData {
		smsDats = append(smsDats, v)
	}
	return smsDats
}

func SmsData() {
	storageSMS := NewStorageSD()
	providers := []string{"Topol", "Rond", "Kildy"}
	countriesCSV := "../StatusPage/config/country.csv"
	countriesString := helpers.CsvInString(countriesCSV)
	//fmt.Println(countriesString)

	smsDataCSV := "../StatusPage/simulator/sms.data"
	smsDataString := helpers.CsvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 4)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 3)
	splitStrings = helpers.ExaminationCoutry(splitStrings, countriesString)
	//splitStrings = examinationInts(splitStrings)

	for i, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewSMSData(s)
		storageSMS.put(l, i)
	}

	for _, v := range storageSMS.getAll() {
		fmt.Println(v)
	}

}

/*func examinationInts(s []string) []string {
	for i, str := range s {
		s := strings.Split(str, ";")
		_, err := strconv.Atoi(s[1])
		if err != nil {
			s = append(s[:i], s[i+1:]...)
			break
		}
		_, err2 := strconv.Atoi(s[2])
		if err2 != nil {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}*/
