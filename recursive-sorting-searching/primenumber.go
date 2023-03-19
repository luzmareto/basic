package main

import (
	"fmt"
	"math"
)

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number))) //sqrt adalah akar
	for i := 2; i <= sqrtNumber; i++ {            //cara menebak bilangan prima atau bukan
		if number%i == 0 {
			return false
		}
	}
	return true
}

func checkPrimeBiasa(number int) bool {
	if number < 2 {
		return false
	}
	var temp int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			temp++
		}
	}
	if temp > 2 {
		return false
	}
	return true
}

func main() {
	fmt.Println(checkPrime(11)) //true, karna 11 adalah bilangan prima
	fmt.Println(checkPrimeBiasa(11))
}
