package helpers

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func CountryString() string {
	s := "AD AE AF AG AI AL AM AO AQ AR AS AT AU AW AX AZ BA BB BD BE BF BG BH BI BJ BL BM BN BO BQ BR BS BT BV BW" +
		" BY BZ CA CC CD CF CG CH CI CK CL CM CN CO CR CU CV CW CX CY CZ DE DJ DK DM DO DZ EC EE EG EH ER ES ET FI " +
		"FJ FK FM FO FR GA GB GD GE GF GG GH GI GL GM GN GP GQ GR GS GT GU GW GY HK HM HN HR HT HU ID IE IL IM IN IO " +
		"IQ IR IS IT JE JM JO JP KE KG KH KI KM KN KP KR KW KY KZ LA LB LC LI LK LR LS LT LU LV LY MA MC MD ME MF MG MH " +
		"MK ML MM MN MO MP MQ MR MS MT MU MV MW MX MY MZ NA NC NE NF NG NI NL NO NP NR NU NZ OM PA PE PF PG PH PK PL " +
		"PM PN PR PS PT PW PY QA RE RO RS RU RW SA SB SC SD SE SG SH SI SJ SK SL SM SN SO SR SS ST SV SX SY SZ TC TD " +
		"TF TG TH TJ TK TL TM TN TO TR TT TV TW TZ UA UG UM US UY UZ VA VC VE VG VI VN VU WF WS YE YT ZA ZM ZW"
	return s
}

func CsvInString(csv string) string {
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
	//s = strings.Replace(s, "\n", " ", -1)
	return s
}

func ExaminationLen(s []string, k int) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")
		if len(splitValues) != k {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func ExaminationProvaiders(s []string, p []string, g int) []string {
	for i := 0; i < len(s); i++ {
		//k := 0
		splitValues := strings.Split(s[i], ";")
		b := CheckProviders(splitValues[g], p)
		if b == false {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func CheckProviders(s string, p []string) bool {
	b := false
	for _, val := range p {
		if strings.ToUpper(val) == strings.ToUpper(s) {
			b = true
			break
		}
	}
	return b
}

func ExaminationCountry(s []string, p string) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")
		l := CheckCountry(p, splitValues[0])
		if l == true {
			continue
		} else {
			s = append(s[:i], s[i+1:]...)
			i--
		}

	}
	return s
}

func CheckCountry(s string, p string) bool {
	//fmt.Printf("ищем подстроку %s в строке %s \n", s, p)
	l := strings.Contains(strings.ToUpper(s), strings.ToUpper(p))
	return l
}

func StringToint(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func StringToFloat32(s string) float32 {
	i, _ := strconv.ParseFloat(s, 32)
	return float32(i)
}

func SliceByteToSliceString(content []byte) []string {
	stringContent := string(content)
	stringContent = strings.Trim(stringContent, "{}")
	stringSliceContent := strings.Split(stringContent, ",")
	return stringSliceContent
}

/*func putInStorage(slice []string, s *storageSupport, k int) {
	i := 0
	var miniStringSlice []string
	for _, value := range slice {
		miniStringSlice = append(miniStringSlice, value)
		if i == k {
			supportString := "{" + strings.Join(miniStringSlice, ",") + "}"
			support := CreateMMS([]byte(supportString))
			s.put(support)
			miniStringSlice = nil
			i = 0
			continue
		}
		i++
	}
}*/

/*func CreateMMS(b []byte) *SupportData {
	var sd *SupportData
	if err := json.Unmarshal(b, &sd); err != nil {
		sd = nil
	}
	return sd
}
*/
