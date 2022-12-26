package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"strings"

	log "github.com/sirupsen/logrus"
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

func VoiceCall(c chan VoiceData) chan VoiceData {
	log.Info("Получаем данные voice_call")
	providers := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	countriesString := helpers.CountryString()
	var storageVoice VoiceData
	smsDataCSV := "./simulator/voice.data"
	smsDataString, err := helpers.CsvInString(smsDataCSV)
	if err != nil {
		log.Info(err)
		storageVoice.err = err
		c <- storageVoice
		return c
	}
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 8)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 3)
	splitStrings = helpers.ExaminationCountry(splitStrings, countriesString)
	for _, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewVoiceData(s)
		storageVoice.voice = append(storageVoice.voice, *l)
	}
	log.Info("Получены данные voice_call")
	storageVoice.err = err
	c <- storageVoice
	return c
}
