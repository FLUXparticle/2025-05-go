package main

import "C"

//export Add
func Add(a, b C.int) C.int {
	return a + b
}

//export IsPrime
func IsPrime(n C.int) C.int {
	x := int(n)
	if x < 2 {
		return 0
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return 0
		}
	}
	return 1
}

func main() {} // muss bei buildmode=c-shared existieren
