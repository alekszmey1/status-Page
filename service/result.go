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
	sms2 := SliceCountryReplace(smsSlice, cm)
	var newSmsSlice []SMSData
	for _, i := range sms2 {
		newSmsSlice = append(newSmsSlice, *i)
	}
	smsProvider := sortProvider(newSmsSlice)
	smsCountry := sortCountry(newSmsSlice)
	var sliceSliceSms [][]SMSData
	sliceSliceSms = append(sliceSliceSms, smsProvider, smsCountry)
	for _, sm := range sliceSliceSms {
		fmt.Println(sm)
	}
}

func sortProvider(st []SMSData) []SMSData {
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
func sortCountry(st []SMSData) []SMSData {
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
func countryReplace(s *SMSData, m map[string]string) *SMSData {
	for reduction, fullName := range m {
		if reduction == s.Country {
			s.Country = fullName
			break
		}
	}
	return s
}
func SliceCountryReplace(s []*SMSData, m map[string]string) []*SMSData {
	for _, i2 := range s {
		countryReplace(i2, m)
	}
	return s

}
