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

type VoiceData struct {
	voice []VoiceCallData
	err   error
}
type DataBilling struct {
	bil BillingData
	err error
}
type smsData struct {
	sms [][]SMSData
	err error
}
type mmsData struct {
	mms [][]MMSData
	err error
}
type emailData struct {
	email map[string][][]EmailData
	err   error
}
type support struct {
	sup []int
	err error
}
type incidData struct {
	inc []IncidentData
	err error
}

func MakeResultT() ResultT {
	r := ResultT{}
	rst, err := GetResultData()
	if err != nil {
		r.Status = false
		r.Error = "Error on collect data"
		return r
	}
	r.Status = true
	r.Data = rst
	return r
}

func GetResultData() (ResultSetT, error) {
	r := ResultSetT{}
	sd := <-chanSms()
	if sd.err != nil {
		log.Info(sd.err)
		return r, sd.err
	}
	mms := <-chanMms()
	if mms.err != nil {
		log.Info(mms.err)
		return r, mms.err
	}
	mail := <-chanEmail()
	if mail.err != nil {
		log.Info(mail.err)
		return r, mail.err
	}
	bil := <-chanBil()
	if bil.err != nil {
		log.Info(bil.err)
		return r, bil.err
	}
	voice := <-chanVoice()
	if voice.err != nil {
		log.Info(voice.err)
		return r, voice.err
	}
	inc := <-chanInc()
	if inc.err != nil {
		log.Info(inc.err)
		return r, inc.err
	}
	sup := <-chanSup()
	if sup.err != nil {
		log.Info(sup.err)
		return r, sup.err
	}
	r = ResultSetT{
		SMS:       sd.sms,
		MMS:       mms.mms,
		VoiceCall: voice.voice,
		Email:     mail.email,
		Billing:   bil.bil,
		Support:   sup.sup,
		Incidents: inc.inc,
	}
	log.Info("заполнен r = Resultset")
	return r, nil
}

func chanSms() chan smsData {
	c := make(chan smsData)
	go sortSMSOne(c)
	return c
}
func chanMms() chan mmsData {
	c := make(chan mmsData)
	go sortMMSOne(c)
	return c
}
func chanEmail() chan emailData {
	c := make(chan emailData)
	go sortEmail(c)
	return c
}
func chanSup() chan support {
	c := make(chan support)
	go sortSupport(c)
	return c
}
func chanInc() chan incidData {
	c := make(chan incidData)
	go sortIncident(c)
	return c
}
func chanVoice() chan VoiceData {
	c := make(chan VoiceData)
	go VoiceCall(c)
	return c
}
func chanBil() chan DataBilling {
	c := make(chan DataBilling)
	go Billing(c)
	return c
}

func sortSMSOne(c chan smsData) chan smsData {
	log.Info("запущена сортировка смс")
	var sms smsData
	smsSlice, err := SmsData()
	if err != nil {
		log.Infof("ошибка сортировки смс %s", err)
		sms.err = err
		c <- sms
		return c
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

func sortMMSOne(c chan mmsData) chan mmsData {
	log.Info("запущена сортировка ммс")
	var sliceSliceMms mmsData
	smsSlice, err := MmsData()
	if err != nil {
		log.Infof("ошибка сортировки ммс %s", err)
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

func sortEmail(c chan emailData) chan emailData {
	emailMap := emailData{}
	emailSlice, err := Email()
	if err != nil {
		log.Infof("ошибка сортировки email %s", err)
		emailMap.err = err
		c <- emailMap
		return c
	}
	countryEmail := sortCountryEmail(emailSlice)
	m := makeMapEmail(countryEmail)
	emailMap.email = minAndMaxValueMap(m)
	log.Info("Проведена сортировка email")
	c <- emailMap
	return c
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

func sortSupport(c chan support) chan support {
	var sort support
	sup, err := Support()
	if err != nil {
		log.Infof("ошибка сортировки support %s", err)
		sort.err = err
		c <- sort
		return c
	}
	sumTic := sumTickets(sup)
	loading := load(sumTic)
	waitTime := waitingTime(sumTic)
	sort.sup = []int{loading, waitTime}
	log.Info("Проведена сортировка support")
	c <- sort
	return c
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

func sortIncident(c chan incidData) chan incidData {
	var i incidData
	inc, err := Incident()
	if err != nil {
		i.err = err
		log.Infof("ошибка сортировки incident %s", err)
		c <- i
		return c
	}
	i.inc = sortingIncident(inc)
	log.Info("Проведена сортировка incident")
	c <- i
	return c
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
