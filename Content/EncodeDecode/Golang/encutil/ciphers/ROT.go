package encutil

import (
	"strings"
)

var _lookupEnc map[string]int
var _lookupDec map[int]string
var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

/*
EncodeROT returns an encoded string using ROT cipher

@param base: determines to what extent the alphabets must be rotated.

@param direction: determines in which direction the letters to be rotated.
*/
func EncodeROT(data string, base int, direction rune) (res string) {
	_tempLookup, _ := _makeMaps(base, direction)

	for _, elem := range data {
		res += _tempLookup[string(elem)]
	}

	return
}

/*
DecodeROT returns an decoded string which was encoded using ROT cipher

@param base: determines to what extent the alphabets must be rotated
*/
func DecodeROT(data string, base int, direction rune) (res string) {
	_, _revLookup := _makeMaps(base, direction)

	for _, elem := range data {
		res += _revLookup[string(elem)]
	}
	return
}

func rotL(m int, arr []string) []string {
	newArr := make([]string, len(arr))

	for i, k := range arr {
		newPos := (((i - m) % len(arr)) + len(arr)) % len(arr)
		newArr[newPos] = k
	}

	return newArr
}

func rotR(m int, arr []string) []string {

	newArr := make([]string, len(arr))

	for i, k := range arr {
		newPos := (i + m) % len(arr)
		newArr[newPos] = k
	}
	return newArr
}

func _makeMaps(base int, direction rune) (map[string]string, map[string]string) {
	_tempLookup := make(map[string]string)
	_revMap := make(map[string]string)
	var rotStr []string

	l := strings.Split(letters, "")

	if direction == 'l' || direction == 'L' {
		rotStr = rotL(base, l)
	} else if direction == 'r' || direction == 'R' {
		rotStr = rotR(base, l)
	}

	for i, c := range rotStr {
		_tempLookup[c] = string(letters[i])
		_revMap[string(letters[i])] = c
	}

	for _, i := range _symbols {
		_tempLookup[string(i)] = string(i)
		_revMap[string(i)] = string(i)
	}

	return _tempLookup, _revMap
}
