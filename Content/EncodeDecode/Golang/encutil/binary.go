package encutil

import (
	"errors"
	"strconv"
)

/*
ToBin converts data into respective binary representation of it

mainly uses the ***strconv*** lib: https://golang.org/pkg/strconv/#pkg-overview
*/
func ToBin(data ...interface{}) (binstr string) {
	return
}

/*
BinaryToString converts a string with binary into a decoded string

*/
func BinaryToString(binStr string) (string, error) {

	if len(binStr)%8 != 0 {
		return "", errors.New("Invalid Binary String")
	}

	return "", nil
}

func _intToBin(data int) string {
	return strconv.FormatInt(int64(data), 2)
}

func _int64ToBin(data int64) string {
	return strconv.FormatInt(data, 2)
}

func _uintToBin(data uint) string {
	return strconv.FormatUint(uint64(data), 2)
}

func _uint64ToBin(data uint64) string {
	return strconv.FormatUint(data, 2)
}

func _float32ToBin(data float32, precision int) string {
	return strconv.FormatFloat(float64(data), 'b', precision, 2)
}

func _float64ToBin(data float64, precision int) string {
	return strconv.FormatFloat(data, 'b', precision, 2)
}

func _runeToBin(data rune) string {
	return ""
}

func _strToBin(data string) string {
	// str := strings.Split(data, "")
	if data == "" {
		return "00000000"
	}

	return ""
}
