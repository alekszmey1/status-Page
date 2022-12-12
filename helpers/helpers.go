package helpers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
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
func CountryMap() map[string]string {

	m := make(map[string]string)
	{
		m["AD"] = "Andorra"
		m["AE"] = "United Arab Emirates"
		m["AF"] = "Afghanistan"
		m["AG"] = "Antigua and Barbuda"
		m["AI"] = "Anguilla"
		m["AL"] = "Albania"
		m["AM"] = "Armenia"
		m["AO"] = "Angola"
		m["AR"] = "Argentina"
		m["AS"] = "American Samoa"
		m["AT"] = "Austria"
		m["AU"] = "Australia"
		m["AW"] = "Aruba"
		m["AX"] = "Åland Islands"
		m["AZ"] = "Azerbaijan"
		m["BA"] = "Bosnia and Herzegovina"
		m["BB"] = "Barbados"
		m["BD"] = "Bangladesh"
		m["BE"] = "Belgium"
		m["BF"] = "Burkina Faso"
		m["BG"] = "Bulgaria"
		m["BH"] = "Bahrain"
		m["BI"] = "Burundi"
		m["BJ"] = "Benin"
		m["BL"] = "Saint Barthelemy"
		m["BM"] = "Bermuda"
		m["BN"] = "Brunei Darussalam"
		m["BO"] = "Bolivia"
		m["BQ"] = "Bonaire"
		m["BR"] = "Brazil"
		m["BS"] = "Bahamas"
		m["BT"] = "Bhutan"
		m["BV"] = "Bouvet Island"
		m["BW"] = "Botswana"
		m["BY"] = "Belarus"
		m["BZ"] = "Belize"
		m["CA"] = "Canada"
		m["CC"] = "Cocos"
		m["CD"] = "Congo (the Democratic Republic of the)"
		m["CF"] = "Central African Republic"
		m["CG"] = "Congo"
		m["CH"] = "Switzerland"
		m["CI"] = "Côte d'Ivoire"
		m["CK"] = "Cook Islands (the)"
		m["CL"] = "Chile"
		m["CM"] = "Cameroon"
		m["CN"] = "China"
		m["CO"] = "Colombia"
		m["CR"] = "Costa Rica"
		m["CU"] = "Cuba"
		m["CV"] = "Cabo Verde"
		m["CW"] = "Curaçao"
		m["CX"] = "Christmas Island"
		m["CY"] = "Cyprus"
		m["CZ"] = "Czechia"
		m["DE"] = "Germany"
		m["DJ"] = "Djibouti"
		m["DK"] = "Denmark"
		m["DM"] = "Dominica"
		m["DO"] = "Dominican Republic"
		m["DZ"] = "Algeria"
		m["EC"] = "Ecuador"
		m["EE"] = "Estonia"
		m["EG"] = "Egypt"
		m["EH"] = "Western Sahara"
		m["ER"] = "Eritrea"
		m["ES"] = "Spain"
		m["ET"] = "Ethiopia"
		m["FI"] = "Finland"
		m["FJ"] = "Fiji"
		m["FK"] = "Falkland Islands"
		m["FM"] = "Micronesia"
		m["FO"] = "Faroe Islands"
		m["FR"] = "France"
		m["GA"] = "Gabon"
		m["GB"] = "United Kingdom of Great Britain and Northern Ireland (the)"
		m["GD"] = "Grenada"
		m["GE"] = "Georgia"
		m["GF"] = "French Guiana"
		m["GG"] = "Guernsey"
		m["GH"] = "Ghana"
		m["GI"] = "Gibraltar"
		m["GL"] = "Greenland"
		m["GM"] = "Gambia"
		m["GN"] = "Guinea"
		m["GP"] = "Guadeloupe"
		m["GQ"] = "Equatorial Guinea"
		m["GR"] = "Greece"
		m["GS"] = "South Georgia and the South Sandwich Islands"
		m["GT"] = "Guatemala"
		m["GU"] = "Guam"
		m["GW"] = "Guinea-Bissau"
		m["GY"] = "Guyana"
		m["HK"] = "Hong Kong"
		m["HM"] = "Heard Island and McDonald Islands"
		m["HN"] = "Honduras"
		m["HR"] = "Croatia"
		m["HT"] = "Haiti"
		m["HU"] = "Hungary"
		m["ID"] = "Indonesia"
		m["IE"] = "Ireland"
		m["IL"] = "Israel"
		m["IM"] = "Isle of Man"
		m["IN"] = "India"
		m["IQ"] = "Iraq"
		m["IR"] = "Iran"
		m["IS"] = "Iceland"
		m["IT"] = "Italy"
		m["JE"] = "Jersey"
		m["JM"] = "Jamaica"
		m["JO"] = "Jordan"
		m["JP"] = "Japan"
		m["KE"] = "Kenya"
		m["KG"] = "Kyrgyzstan"
		m["KH"] = "Cambodia"
		m["KI"] = "Kiribati"
		m["KM"] = "Comoros"
		m["KN"] = "Saint Kitts and Nevis"
		m["KP"] = "Korea (the Democratic People's Republic of)"
		m["KR"] = "Korea (the Republic of)"
		m["KW"] = "Kuwait"
		m["KY"] = "Cayman Islands (the)"
		m["KZ"] = "Kazakhstan"
		m["LA"] = "Lao People's Democratic Republic (the)"
		m["LB"] = "Lebanon"
		m["LC"] = "Saint Lucia"
		m["LI"] = "Liechtenstein"
		m["LK"] = "Sri Lanka"
		m["LR"] = "Liberia"
		m["LS"] = "Lesotho"
		m["LT"] = "Lithuania"
		m["LU"] = "Luxembourg"
		m["LV"] = "Latvia"
		m["LY"] = "Libya"
		m["MA"] = "Morocco"
		m["MC"] = "Monaco"
		m["MD"] = "Moldova (the Republic of)"
		m["ME"] = "Montenegro"
		m["MF"] = "Saint Martin"
		m["MG"] = "Madagascar"
		m["MH"] = "Marshall Islands (the)"
		m["MK"] = "North Macedonia"
		m["ML"] = "Mali"
		m["MM"] = "Myanmar"
		m["MN"] = "Mongolia"
		m["MO"] = "Macao"
		m["MP"] = "Northern Mariana Islands"
		m["MQ"] = "Martinique"
		m["MR"] = "Mauritania"
		m["MS"] = "Montserrat"
		m["MT"] = "Malta"
		m["MU"] = "Mauritius"
		m["MV"] = "Maldives"
		m["MW"] = "Malawi"
		m["MX"] = "Mexico"
		m["MY"] = "Malaysia"
		m["MZ"] = "Mozambique"
		m["NA"] = "Namibia"
		m["NC"] = "New Caledonia"
		m["NE"] = "Niger"
		m["NF"] = "Norfolk Island"
		m["NG"] = "Nigeria"
		m["NI"] = "Nicaragua"
		m["NL"] = "Netherlands"
		m["NO"] = "Norway"
		m["NP"] = "Nepal"
		m["NR"] = "Nauru"
		m["NU"] = "Niue"
		m["NZ"] = "New Zealand"
		m["OM"] = "Oman"
		m["PA"] = "Panama"
		m["PE"] = "Peru"
		m["PF"] = "French Polynesia"
		m["PG"] = "Papua New Guinea"
		m["PH"] = "Philippines"
		m["PK"] = "Pakistan"
		m["PL"] = "Poland"
		m["PM"] = "Saint Pierre and Miquelon"
		m["PN"] = "Pitcairn"
		m["PR"] = "Puerto Rico"
		m["PS"] = "Palestine"
		m["PT"] = "Portugal"
		m["PW"] = "Palau"
		m["PY"] = "Paraguay"
		m["QA"] = "Qatar"
		m["RE"] = "Réunion"
		m["RO"] = "Romania"
		m["RS"] = "Serbia"
		m["RU"] = "Russian Federation"
		m["RW"] = "Rwanda"
		m["SA"] = "Saudi Arabia"
		m["SB"] = "Solomon Islands"
		m["SC"] = "Seychelles"
		m["SD"] = "Sudan"
		m["SE"] = "Sweden"
		m["SG"] = "Singapore"
		m["SH"] = "Saint Helena, Ascension and Tristan da Cunha"
		m["SI"] = "Slovenia"
		m["SJ"] = "Svalbard"
		m["SK"] = "Slovakia"
		m["SL"] = "Sierra Leone"
		m["SM"] = "San Marino"
		m["SN"] = "Senegal"
		m["SO"] = "Somalia"
		m["SR"] = "Suriname"
		m["SS"] = "South Sudan"
		m["ST"] = "Sao Tome and Principe"
		m["SV"] = "El Salvador"
		m["SX"] = "Sint Maarten"
		m["SY"] = "Syrian Arab Republic"
		m["SZ"] = "Eswatini"
		m["TC"] = "Turks and Caicos Islands"
		m["TD"] = "Chad"
		m["TF"] = "French Southern Territories (the)"
		m["TG"] = "Togo"
		m["TH"] = "Thailand"
		m["TJ"] = "The Republic of Tajikistan"
		m["TK"] = "Tokelau"
		m["TL"] = "Timor-Leste"
		m["TM"] = "Turkmenistan"
		m["TN"] = "The Republic of Tunisia"
		m["TO"] = "Tonga"
		m["TR"] = "The Republic of Türkiye"
		m["TT"] = "The Republic of Trinidad and Tobago"
		m["TV"] = "Tuvalu"
		m["TW"] = "Taiwan (Province of China)"
		m["TZ"] = "Tanzania, the United Republic of"
		m["UA"] = "Ukraine"
		m["UG"] = "Uganda"
		m["UM"] = "United States Minor Outlying Islands"
		m["US"] = "United States of America"
		m["UY"] = "Uruguay"
		m["UZ"] = "Uzbekistan"
		m["VA"] = "Holy See"
		m["VC"] = "Saint Vincent and the Grenadines"
		m["VE"] = "Venezuela (Bolivarian Republic of)"
		m["VG"] = "Virgin Islands (British)"
		m["VI"] = "Virgin Islands (U.S.)"
		m["VN"] = "Viet Nam"
		m["VU"] = "Vanuatu"
		m["WF"] = "Wallis and Futuna"
		m["WS"] = "Samoa"
		m["YE"] = "Yemen"
		m["YT"] = "Mayotte"
		m["ZA"] = "South Africa"
		m["ZM"] = "Zambia"
		m["ZW"] = "Zimbabwe"
	}
	return m
}

func CsvInString(csv string) (string, error) {
	file, err := os.Open(csv)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	s := string(bytes)
	return s, err
}

func ExaminationLen(s []string, k int) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")
		if len(splitValues) != k {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	log.Info("убрали данные, не соответствующие длине")
	return s

}

func ExaminationProvaiders(s []string, p []string, g int) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")
		b := CheckProviders(splitValues[g], p)
		if b == false {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	log.Info("убрали данные, не соответствующие названиям провайдеров")
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
	log.Info("убрали данные, не соответствующие названию стран")
	return s
}

func CheckCountry(s string, p string) bool {

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

func StringToSliceString(s string) []string {
	s2 := strings.Trim(s, "[][]")
	s2 = strings.Replace(s2, "[", "", -1)
	s = strings.Replace(s2, "},{", "};{", -1)
	str := strings.Split(s, ";")
	log.Info("убрали лишние скобки, разбили строку на массив строк")
	return str
}

func UrlToString(url string) (string, error) {
	var s string
	resp, err := http.Get(url)
	if err != nil {
		return s, fmt.Errorf("получение данных с url  %s выдало ошибку %s \n", url, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return s, fmt.Errorf("support respose failed with status code %d : \n", resp.StatusCode)
	}
	bufer, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}
	s = string(bufer)
	log.Infof("получили данные с url %s", url)
	return s, nil
}
