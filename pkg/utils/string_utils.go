package utils

import (
	"bytes"
	"encoding/json"
	"strings"
	"unicode/utf8"
)

func TruncateEnd(str string, length int, omission string) (string, int) {
	r := []rune(str)
	sLen := len(r)
	if length >= sLen {
		return str, sLen
	}

	return string(string(r[:length-utf8.RuneCountInString(omission)]) + omission), length
}

func ArrToString(str []string, maxLen int) (string,int) {

	fullString := strings.Join(str[:], ",")

	if maxLen == -1 {
		return fullString, utf8.RuneCountInString(fullString)
	}

	return TruncateEnd(fullString, maxLen,"...")
}

func PrettyJSONString(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}