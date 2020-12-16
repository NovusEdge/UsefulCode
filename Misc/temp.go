package usefulcode

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

//IsPrime :Checks primality of a positive integer
func IsPrime(n int) (bool, error) {
	if n < 0 {
		return false, errors.New("Number negative")
	} else if n == 2 {
		return true, nil
	}
	upper := int(math.Trunc(math.Sqrt(float64(n))))
	for i := 2; i < upper; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

//IsPallindrome :Checks if a positive integer is a pallindrome
func IsPallindrome(n int) bool {
	s := fmt.Sprintf("%d", int(math.Abs(float64(n))))
	return s == _reverse(s)
}

func _reverse(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) { ret += string(r) }(v)
	}
	return
}

//Factors : Returns a list of all factors of a number
func Factors(n int) (res []int) {
	upper := int((math.Sqrt(math.Abs(float64(n)))))
	for i := 1; i < upper; i++ {
		if n%i == 0 {
			res = append(res, i)
			res = append(res, int(n/i))
		}
	}
	return
}

/*
DigitalRoot : Returns the sum of all digits of a number
* @param [n]: Number for which the digital root is calculated
* @param [downToOne]: Pass true if wanting to trim the number down to a one digit number
If [downToOne] is passed as true: (EXAMPLE) DigitalRoot(1987, true) => 1+9+8+7 = 25 -> 2+5 = 7
*/
func DigitalRoot(n int, downToOne bool) (res int) {
	if downToOne {
		return _recurDigital(n)
	}
	return _addDigits(n)
}

func _recurDigital(n int) int {
	if n/10 == 0 {
		return n
	}
	return _recurDigital(_addDigits(n))
}

func _addDigits(n int) (res int) {
	s := strconv.Itoa(n)
	var j int
	for _, i := range s {
		j, _ = strconv.Atoi(string(i))
		res += j
	}
	return
}

//PrimeSieve : Implements the Sieve of Eratosthenes
func PrimeSieve(n int) []int {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if integers[p] == true {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}
	return primes
}

/*
Fibonacci :Reports the nth term in the fibonacci series
* Fibonacci(0) --> 1
* Fibonacci(1) --> 1
* Fibonacci(2) --> 2
* ...
*/
func Fibonacci(n int64) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	limit := *big.NewInt(n)

	for a.Cmp(&limit) < 0 {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}
