package encutil

import (
	"math/rand"
)

var _symbols = "!\"#$%&'()*+,-./0123456789:;<=>?@[\\] ^_`{|}~ "
var affineMap map[string]string

type _affKeyMaps struct {
	a int
	b int
}

/*
EncodeAffine returns an encoded string using affine cipher
*/
func EncodeAffine(data string, key int) (res string) {
	m := int(len(letters) / 2)
	var keys _affKeyMaps

	keys._genkeys(m)

	return
}

/*
DecodeAffine decodes a string encoded using affine cipher
*/
func DecodeAffine(data string) (res string) {

	return
}

func (k *_affKeyMaps) _genkeys(mod int) {
	b := k.b

	r := rand.New(rand.NewSource(55))
	a := r.Intn(mod - 1)

	if !_isRelPrime(a, mod) {
		k._genkeys(mod)
	}

	k.a = a
	k.b = b
}

func _isRelPrime(a, b int) bool {
	return gcd(a, b) == 1
}

func gcd(a int, b int) (res int) {

	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			res = i
		}
	}

	return res
}
