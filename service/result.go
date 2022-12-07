package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

func MakeResultT() ResultT {
	r := ResultT{}
	rst := GetResultData()
	b := true
	if rst.MMS == nil || rst.SMS == nil || rst.Incidents == nil || rst.Support == nil || rst.VoiceCall == nil ||
		rst.Email == nil /* || rst.Billing == nil */ {
		b = false
	}
	if b == true {
		r.Data = rst

	} else {
		r.Error = "Error on collect data"
	}
	r.Status = b
	return r
}

func GetResultData() ResultSetT {
	r := ResultSetT{
		SMS:       sortSMSOne(),
		MMS:       sortMMSOne(),
		VoiceCall: VoiceCall(),
		Email:     sortEmail(),
		Billing:   Billing(),
		Support:   sortSupport(),
		Incidents: sortIncident(),
	}
	return r
}

func sortSMSOne() [][]SMSData {
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
	log.Info("Проведена сортировка sms")
	return sliceSliceSms

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

func sortMMSOne() [][]MMSData {
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
	log.Info("Проведена сортировка mms")
	return sliceSliceMms
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

func sortEmail() map[string][][]EmailData {
	emailSlice := Email()
	countryEmail := sortCountryEmail(emailSlice)
	m := makeMapEmail(countryEmail)
	emailMap := minAndMaxValueMap(m)
	log.Info("Проведена сортировка mms")
	return emailMap

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
	//fmt.Println(" проведена сортировка email по странам")
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
	return desiredMap
}

func sortSupport() []int {
	sup := Support()
	sumTic := sumTickets(sup)
	loading := load(sumTic)
	waitTime := waitingTime(sumTic)
	sort := []int{loading, waitTime}
	log.Info("Проведена сортировка support")
	return sort
}

func load(i int) int {
	a := 0
	if i < 9 {
		a = 1
	} else if i > 16 {
		a = 3
	} else {
		a = 2
	}
	return a
}
func sumTickets(s []SupportData) int {
	x := 0
	for _, data := range s {
		x += data.ActiveTickets
	}
	return x
}
func waitingTime(i int) int {
	x := i * (60 / 18)
	return x
}

func sortIncident() []IncidentData {
	inc := Incident()
	incSort := sortingIncident(inc)
	log.Info("Проведена сортировка incident")
	return incSort
}

func sortingIncident(st []IncidentData) []IncidentData {
	s := make([]IncidentData, len(st))
	copy(s, st[:])
	for i := 0; i < len(s); i++ {
		var y = i
		for j := i; j < len(s); j++ {
			if strings.Compare(s[i].Status, s[j].Status) > 0 {
				y = j
				s[i], s[y] = s[y], s[i]
			}
		}
	}

	return s
}
