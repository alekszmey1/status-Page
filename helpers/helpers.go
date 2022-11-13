package helpers

import (
	"io"
	"os"
	"strings"
)

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

func ExaminationProvaiders(s []string, p []string) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")
		if splitValues[3] != p[0] || splitValues[3] != p[1] || splitValues[3] != p[2] {
			continue
		} else {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func ExaminationCoutry(s []string, p string) []string {
	for i := 0; i < len(s); i++ {
		splitValues := strings.Split(s[i], ";")

		l := strings.Contains(strings.ToUpper(p), strings.ToUpper(splitValues[0]))
		if l == true {
			continue
		} else {
			s = append(s[:i], s[i+1:]...)
			i--
		}

	}
	return s
}
