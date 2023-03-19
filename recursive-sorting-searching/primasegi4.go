package main

import (
	"fmt"
	"math"
)

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func primeAfterNumber(number int) int {
	number++
	for !checkPrime(number) {
		number++
	}
	return number
}

func primaSegiEmpat(wide, high, start int) string {
	var jumlah int
	var result string
	var number = start
	for i := 1; i <= high; i++ {
		for j := 1; j <= wide; i++ {
			//proses mendapatkan bilangan prima setelah start
			number = primeAfterNumber(number)
			result = fmt.Sprintf("%v\t", number)
			jumlah += number
		}
		result += "\n"
	}
	result = fmt.Sprintf("%v\t", jumlah)
	return result
}

func main() {
	fmt.Println(primaSegiEmpat(2, 3, 13))
}
