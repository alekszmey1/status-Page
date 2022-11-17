package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"fmt"
	"strings"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func NewEmailData(str []string) *EmailData {
	ed := EmailData{}
	ed.Country = str[0]
	ed.Provider = str[1]
	ed.DeliveryTime = helpers.StringToint(str[2])
	return &ed
}

type StorageED struct {
	storageEmailData map[int]*EmailData
}

func NewStorageED() *StorageED {
	return &StorageED{storageEmailData: make(map[int]*EmailData)}
}
func (u *StorageED) put(vd *EmailData, i int) {
	u.storageEmailData[i] = vd
}
func (u *StorageED) getAll() []*EmailData {
	var voiceDats []*EmailData
	for _, v := range u.storageEmailData {
		voiceDats = append(voiceDats, v)
	}
	return voiceDats
}

func Email() {

	providers := []string{"Orange", "Comcast", "AOL", "Gmail", "Yahoo", "Hotmail", "MSN", "Live", "RediffMail", "GMX",
		"Protonmail", "Yandex", "Mail.ru"}
	countriesString := helpers.CountryString()

	storageVD := NewStorageED()
	smsDataCSV := "../StatusPage//simulator/email.data"
	smsDataString := helpers.CsvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 3)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 1)
	splitStrings = helpers.ExaminationCoutry(splitStrings, countriesString)

	for i, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewEmailData(s)
		storageVD.put(l, i)
	}

	for _, v := range storageVD.getAll() {
		fmt.Println(v)
	}

}
