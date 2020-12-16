package encutil

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

/*
ToBinary converts data into respective binary representation of it

mainly uses the ***strconv*** lib: https://golang.org/pkg/strconv/#pkg-overview

* @param prec: for precision in converting floatpoint numbers to binary

* @param sep:  for defining what seperator is to be used to parse a string for conversion
*/
func ToBinary(data interface{}, prec int, sep string) string {

	switch v := reflect.ValueOf(data); v.Kind() {
	case reflect.Int32:
		return _intToBin(int(v.Int()))

	case reflect.Int64:
		return _int64ToBin(v.Int())

	case reflect.Uint32:
		return _uintToBin(uint(v.Uint()))

	case reflect.Uint64:
		return _uint64ToBin(v.Uint())

	case reflect.Float32:
		return _float32ToBin(float32(v.Float()), prec)

	case reflect.Float64:
		return _float64ToBin(v.Float(), prec)

	case reflect.String:
		return _strToBin(v.String(), sep)
	}

	return ""
}

/*
BinaryToString converts a string with binary into a decoded string

*/
func BinaryToString(binStr string) (string, error) {

	if len(binStr)%8 != 0 {
		return "", errors.New("Invalid Binary String")
	}

	result := ""

	for i := 0; i < len(binStr)-8; i += 8 {
		temp, _ := strconv.ParseInt(binStr[i:i+8], 10, 32)
		result += string(int32(temp))
	}

	return result, nil
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

func _strToBin(data string, sep string) string {
	_runeToBin := func(data rune) string {
		return strconv.FormatInt(int64(data), 2)
	}

	if data == "" {
		return "00000000"
	}

	str := strings.Split(data, sep)
	binStr := ""
	for _, i := range str {
		for _, j := range i {
			binStr += _runeToBin(j)
		}
	}

	return binStr
}
