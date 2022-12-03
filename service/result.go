package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"fmt"
	"strings"
)

/*
type Result struct {
	Status bool       `json:"status"`
	Data   ResultSelt `json:"data"`
	Error  string     `json:"error"`
}


type ResultSelt struct {
	SMS         [][]SMSData              `json:"sms"`
	MMS         [][]MMSData              `json:"mms"`
	VoiceCall   []VoiceCallData          `json:"voice_call"`
	Email       map[string][][]EmailData `json:"email"`
	BillingData `json:"billing"`
	Support     []int          `json:"support"`
	Incident    []IncidentData `json:"incident"`
}*/

func SortSMSOne() {
	smsSlice := SmsData()
	cm := helpers.CountryMap()
	sms2 := sliceCountryReplaceSMS(smsSlice, cm)
	var newSmsSlice []SMSData
	for _, i := range sms2 {
		newSmsSlice = append(newSmsSlice, *i)
	}
	smsProvider := sortProviderSMS(newSmsSlice)
	smsCountry := sortCountrySMS(newSmsSlice)
	var sliceSliceSms [][]SMSData
	sliceSliceSms = append(sliceSliceSms, smsProvider, smsCountry)
}

func sortProviderSMS(st []SMSData) []SMSData {
	s := make([]SMSData, len(st))
	copy(s, st[:])
	for i := 0; i < len(s); i++ {
		var y = i
		for j := i; j < len(s); j++ {
			if strings.Compare(s[i].Provider, s[j].Provider) > 0 {
				y = j
				s[i], s[y] = s[y], s[i]
			}
		}
	}
	return s
}
func sortCountrySMS(st []SMSData) []SMSData {
	s := make([]SMSData, len(st))
	copy(s, st[:])
	for i := 0; i < len(s); i++ {
		var y = i
		for j := i; j < len(s); j++ {
			if strings.Compare(s[i].Country, s[j].Country) > 0 {
				y = j
				s[i], s[y] = s[y], s[i]
			}
		}
	}
	return s
}
func countryReplaceSMS(s *SMSData, m map[string]string) *SMSData {
	for reduction, fullName := range m {
		if reduction == s.Country {
			s.Country = fullName
			break
		}
	}
	return s
}
func sliceCountryReplaceSMS(s []*SMSData, m map[string]string) []*SMSData {
	for _, i2 := range s {
		countryReplaceSMS(i2, m)
	}
	return s
}

func SortMMSOne() {
	smsSlice := MmsData()
	cm := helpers.CountryMap()
	sms2 := sliceCountryReplaceMMS(smsSlice, cm)
	var newSmsSlice []MMSData
	for _, i := range sms2 {
		newSmsSlice = append(newSmsSlice, *i)
	}
	mmsProvider := sortProviderMMS(newSmsSlice)
	mmsCountry := sortCountryMMS(newSmsSlice)
	var sliceSliceMms [][]MMSData
	sliceSliceMms = append(sliceSliceMms, mmsProvider, mmsCountry)
	for _, sm := range sliceSliceMms {
		fmt.Println(sm)
	}
}

func sortProviderMMS(st []MMSData) []MMSData {
	s := make([]MMSData, len(st))
	copy(s, st[:])
	for i := 0; i < len(s); i++ {
		var y = i
		for j := i; j < len(s); j++ {
			if strings.Compare(s[i].Provider, s[j].Provider) > 0 {
				y = j
				s[i], s[y] = s[y], s[i]
			}
		}
	}
	return s
}
func sortCountryMMS(st []MMSData) []MMSData {
	s := make([]MMSData, len(st))
	copy(s, st[:])
	for i := 0; i < len(s); i++ {
		var y = i
		for j := i; j < len(s); j++ {
			if strings.Compare(s[i].Country, s[j].Country) > 0 {
				y = j
				s[i], s[y] = s[y], s[i]
			}
		}
	}
	return s
}
func countryReplaceMMS(s *MMSData, m map[string]string) *MMSData {
	for reduction, fullName := range m {
		if reduction == s.Country {
			s.Country = fullName
			break
		}
	}
	return s
}
func sliceCountryReplaceMMS(s []*MMSData, m map[string]string) []*MMSData {
	for _, i2 := range s {
		countryReplaceMMS(i2, m)
	}
	return s
}

func SortEmail() {
	emailSlice := Email()
	countryEmail := sortCountryEmail(emailSlice)
	m := makeMapEmail(countryEmail)
	/*for s, data := range m {
		fmt.Println(s, data)
	}*/
	emailMap := minAndMaxValueMap(m)
	for s, i := range emailMap {
		fmt.Println(s, i)
	}

}
func sortCountryEmail(st []EmailData) []EmailData {
	s := make([]EmailData, len(st))
	copy(s, st[:])
	for i := 0; i < len(s); i++ {
		var y = i
		for j := i; j < len(s); j++ {
			if strings.Compare(s[i].Country, s[j].Country) > 0 {
				y = j
				s[i], s[y] = s[y], s[i]
			}
		}
	}
	fmt.Println(" проведена сортировка email по странам")
	return s
}

func makeMapEmail(e []EmailData) map[string][]EmailData {
	m := make(map[string][]EmailData)
	country := e[0].Country
	var s []EmailData
	for i := 0; i < len(e); i++ {
		if e[i].Country == country {
			s = append(s, e[i])
			if i == len(e)-1 {
				m[country] = s
			}
		} else {
			m[country] = s
			country = e[i].Country
			s = nil
			s = append(s, e[i])
			if i == len(e)-1 {
				m[country] = s
			}
		}
	}
	fmt.Println("сделана сохранение в map по странам")
	return m
}

func minAndMaxValue(st []EmailData) [][]EmailData {
	s := make([]EmailData, len(st))
	copy(s, st[:])
	for i := 0; i < len(s); i++ {
		var y = i
		for j := i; j < len(s); j++ {
			if s[i].DeliveryTime > s[j].DeliveryTime {
				y = j
				s[i], s[y] = s[y], s[i]
			}
		}
	}
	sliceMin := []EmailData{s[0], s[1], s[2]}
	sliceMax := []EmailData{s[len(s)-3], s[len(s)-2], s[len(s)-1]}
	var sliceMinMax [][]EmailData
	sliceMinMax = append(sliceMinMax, sliceMin, sliceMax)
	return sliceMinMax
}

func minAndMaxValueMap(m map[string][]EmailData) map[string][][]EmailData {
	desiredMap := make(map[string][][]EmailData)
	for s, data := range m {
		x := minAndMaxValue(data)
		desiredMap[s] = x
	}
	fmt.Println("проведена сортировка по минимальным и максимальным значениям скорости провайдеров")
	return desiredMap
}
