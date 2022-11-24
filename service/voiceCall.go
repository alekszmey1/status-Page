package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"fmt"
	"strings"
)

type VoiceCallData struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
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

type StorageVD struct {
	storageVoiceData map[int]*VoiceCallData
}

func NewStorageVD() *StorageVD {
	return &StorageVD{storageVoiceData: make(map[int]*VoiceCallData)}
}
func (u *StorageVD) put(vd *VoiceCallData, i int) {
	u.storageVoiceData[i] = vd
}
func (u *StorageVD) getAll() []*VoiceCallData {
	var voiceDats []*VoiceCallData
	for _, v := range u.storageVoiceData {
		voiceDats = append(voiceDats, v)
	}
	return voiceDats
}

func VoiceCall() {

	providers := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	countriesString := helpers.CountryString()

	storageVD := NewStorageVD()
	smsDataCSV := "../StatusPage//simulator/voice.data"
	smsDataString := helpers.CsvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 8)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers, 3)
	splitStrings = helpers.ExaminationCountry(splitStrings, countriesString)

	for i, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewVoiceData(s)
		storageVD.put(l, i)
	}

	for _, v := range storageVD.getAll() {
		fmt.Println(v)
	}

}
