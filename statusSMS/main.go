package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

func main() {
	storageSMS := NewStorageSD()
	providers := []string{"Topol", "Rond", "Kildy"}
	countriesCSV := "country.csv"
	countriesString := csvInString(countriesCSV)
	fmt.Println(countriesString)

	smsDataCSV := "../../simulator/skillbox-diploma/sms.data"
	smsDataString := csvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = examinationLen(splitStrings, 4)
	splitStrings = examinationProvaiders(splitStrings, providers)
	splitStrings = examinationCoutry(splitStrings, countriesString)
	splitStrings = examinationInts(splitStrings)

	for i, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewSMSData(s)
		storageSMS.put(l, i)
	}

	for _, v := range storageSMS.getAll() {
		fmt.Println(v)
	}

}

func examinationLen(s []string, k int) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")
		if len(splitValues) != k {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func examinationProvaiders(s []string, p []string) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")
		if splitValues[3] != p[0] || splitValues[3] != p[1] || splitValues[3] != p[2] {
			continue
		} else {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func examinationCoutry(s []string, p string) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")

		l := strings.Contains(strings.ToUpper(p), strings.ToUpper(splitValues[0]))
		if l == true {
			continue
		} else {
			s = append(s[:i], s[i+1:]...)
			i--
		}

	}
	return s
}

func examinationInts(s []string) []string {
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
}

func csvInString(csv string) string {
	file, err := os.Open(csv)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	s := string(bytes)
	return s
}
