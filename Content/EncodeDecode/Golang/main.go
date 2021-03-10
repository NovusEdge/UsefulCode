package main

import (
	b "encutil/encutil"
	ciphers "encutil/encutil/ciphers"
	"fmt"
)

func main() {
	bin := b.ToBinary("100", 0, ", ")
	fmt.Println(bin)

	rotEnc := ciphers.EncodeROT("lorem34", 10, 'r')
	fmt.Println(rotEnc)
	fmt.Println(ciphers.DecodeROT(rotEnc, 10, 'r'))

	atbEnc := ciphers.EncodeAtbash("HELLO WORLD")
	fmt.Println(atbEnc)
	fmt.Println(ciphers.DecodeAtbash(atbEnc))

	// affEnc := ciphers.EncodeAffine()
	// fmt.Println(affEnc)
	// fmt.Println(ciphers.DecodeAffine())

}
