package encutil

var _subkey map[string]string

/*
EncodeAtbash returns an encoded string using Atbash cipher
*/
func EncodeAtbash(data string) (res string) {
	res = ""
	_makeKey()
	for i := 0; i < len(data); i++ {
		res += _subkey[string(data[i])]
	}
	return
}

/*
DecodeAtbash decodes a string encoded using atbash cipher
*/
func DecodeAtbash(data string) (res string) {
	for i := 0; i < len(data); i++ {
		res += _subkey[string(data[i])]
	}

	return
}

func _makeKey() {
	temp1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	temp2 := "abcdefghijklmnopqrstuvwxyz"
	_res := make(map[string]string)

	for i := 0; i < len(temp1); i++ {
		_res[string(temp1[i])] = string(temp1[len(temp1)-i-1])
		_res[string(temp2[i])] = string(temp2[len(temp2)-i-1])
	}

	for _, i := range _symbols {
		_res[string(i)] = string(i)
	}

	_subkey = _res
}
