package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"fmt"
	"strconv"
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
	vd.ConnectionStability = stringInFloat32(str[4])
	vd.TTFB = stringInInt(str[5])
	vd.VoicePurity = stringInInt(str[6])
	vd.MedianOfCallsTime = stringInInt(str[7])

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
	countriesCSV := "../StatusPage/config/country.csv"
	countriesString := helpers.CsvInString(countriesCSV)

	storageVD := NewStorageVD()
	smsDataCSV := "../StatusPage//simulator/voice.data"
	smsDataString := helpers.CsvInString(smsDataCSV)
	splitStrings := strings.Split(smsDataString, "\n")
	splitStrings = helpers.ExaminationLen(splitStrings, 8)
	splitStrings = helpers.ExaminationProvaiders(splitStrings, providers)
	splitStrings = helpers.ExaminationCoutry(splitStrings, countriesString)
	//splitStrings = examinationInts(splitStrings)

	for i, str := range splitStrings {
		s := strings.Split(str, ";")
		l := NewVoiceData(s)
		storageVD.put(l, i)
	}

	for _, v := range storageVD.getAll() {
		fmt.Println(v)
	}

}

/*func examinationInts(s []string) []string {
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
}*/

func stringInInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func stringInFloat32(s string) float32 {
	i, _ := strconv.ParseFloat(s, 32)
	return float32(i)
}
