package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func NewVoiceData(str []string) *VoiceCallData {
	vd := VoiceCallData{}
	vd.Country = str[0]
	vd.Bandwidth = str[1]
	vd.ResponseTime = str[2]
	vd.Provider = str[3]
	vd.ConnectionStability = helpers.StringToFloat32(str[4])
	vd.TTFB = helpers.StringToint(str[5])
	vd.VoicePurity = helpers.StringToint(str[6])
	vd.MedianOfCallsTime = helpers.StringToint(str[7])

	return &vd
}

func VoiceCall() []VoiceCallData {

	providers := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	countriesString := helpers.CountryString()

	var storageVoice []VoiceCallData
	smsDataCSV := "../StatusPage//simulator/voice.data"
	smsDataString := helpers.CsvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 8)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 3)
	splitStrings = helpers.ExaminationCountry(splitStrings, countriesString)

	for _, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewVoiceData(s)
		storageVoice = append(storageVoice, *l)
	}
	return storageVoice

}
