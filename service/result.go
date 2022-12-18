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

type smsdata struct {
	sms [][]SMSData
	err error
}
type mmsdata struct {
	mms [][]MMSData
	err error
}
type emaildata struct {
	email map[string][][]EmailData
	err   error
}

func MakeResultT() ResultT {
	r := ResultT{}
	rst, err := GetResultData()
	if err != nil {
		r.Status = false
		r.Error = "Error on collect data"
		log.Fatalln(err)
		return r
	}
	r.Status = true
	r.Data = rst
	return r
}

func GetResultData() (ResultSetT, error) {
	r := ResultSetT{}
	smsChan := make(chan smsdata)
	go sortSMSOne(smsChan)
	mmsChan := make(chan mmsdata)
	go sortMMSOne(mmsChan)
	emailChan := make(chan emaildata)
	go sortEmail(emailChan)

	sd := <-smsChan
	if sd.err != nil {
		log.Fatalln(sd.err)
		return r, sd.err
	}

	mms := <-mmsChan
	if mms.err != nil {
		log.Fatalln(mms.err)
		return r, mms.err
	}

	mail := <-emailChan
	if mail.err != nil {
		log.Fatalln(mail.err)
		return r, mail.err
	}
	bil, err := Billing()
	if err != nil {
		log.Fatalln(err)
		return r, err
	}

	voice, err := VoiceCall()
	if err != nil {
		log.Fatalln(err)
		return r, err
	}
	inc, err := sortIncident()
	if err != nil {
		log.Fatalln(err)
		return r, err
	}

	sup, err := sortSupport()
	if err != nil {
		log.Fatalln(err)
		return r, err
	}

	r = ResultSetT{
		SMS:       sd.sms,
		MMS:       mms.mms,
		VoiceCall: voice,
		Email:     mail.email,
		Billing:   bil,
		Support:   sup,
		Incidents: inc,
	}
	return r, err
}

func sortSMSOne(c chan smsdata) chan smsdata {
	log.Info("запущена сортировка смс")
	var sms smsdata
	smsSlice, err := SmsData()
	if err != nil {
		log.Fatalln(err)
		sms.err = err
		c <- sms
	}
	cm := helpers.CountryMap()
	sms2 := sliceCountryReplaceSMS(smsSlice, cm)
	var newSmsSlice []SMSData
	for _, i := range sms2 {
		newSmsSlice = append(newSmsSlice, *i)
	}
	smsProvider := sortProviderSMS(newSmsSlice)
	smsCountry := sortCountrySMS(newSmsSlice)
	sms.sms = append(sms.sms, smsProvider, smsCountry)
	log.Info("Проведена сортировка sms")
	c <- sms
	return c
}

func outChanSms(c chan smsdata) ([][]SMSData, error) {
	var s smsdata
	s = <-c
	return s.sms, s.err
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

func sortMMSOne(c chan mmsdata) chan mmsdata {
	log.Info("запущена сортировка ммс")
	var sliceSliceMms mmsdata
	smsSlice, err := MmsData()
	if err != nil {
		log.Fatalln(err)
		sliceSliceMms.err = err
		c <- sliceSliceMms
		return c
	}
	cm := helpers.CountryMap()
	sms2 := sliceCountryReplaceMMS(smsSlice, cm)
	var newSmsSlice []MMSData
	for _, i := range sms2 {
		newSmsSlice = append(newSmsSlice, *i)
	}
	mmsProvider := sortProviderMMS(newSmsSlice)
	mmsCountry := sortCountryMMS(newSmsSlice)
	sliceSliceMms.mms = append(sliceSliceMms.mms, mmsProvider, mmsCountry)
	log.Info("Проведена сортировка mms")
	c <- sliceSliceMms
	return c
}

func outChanMms(c chan mmsdata) ([][]MMSData, error) {
	var s mmsdata
	s = <-c
	return s.mms, s.err
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

func sortEmail(c chan emaildata) chan emaildata {
	emailMap := emaildata{}
	emailSlice, err := Email()
	if err != nil {
		log.Fatalln(err)
		emailMap.err = err
		c <- emailMap
		return c
	}
	countryEmail := sortCountryEmail(emailSlice)
	m := makeMapEmail(countryEmail)
	emailMap.email = minAndMaxValueMap(m)
	log.Info("Проведена сортировка mms")
	c <- emailMap
	return c
}

func outChanEmail(c chan emaildata) (map[string][][]EmailData, error) {
	var s emaildata
	s = <-c
	return s.email, s.err
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
	log.Info(" проведена сортировка email по странам")
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

func sortSupport() ([]int, error) {
	var sort []int
	sup, err := Support()
	if err != nil {
		log.Fatalln(err)
		return sort, err
	}
	sumTic := sumTickets(sup)
	loading := load(sumTic)
	waitTime := waitingTime(sumTic)
	sort = []int{loading, waitTime}
	log.Info("Проведена сортировка support")
	return sort, err
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

func sortIncident() ([]IncidentData, error) {
	inc, err := Incident()
	if err != nil {
		log.Fatalln(err)
		return inc, err
	}
	incSort := sortingIncident(inc)
	log.Info("Проведена сортировка incident")
	return incSort, err
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
