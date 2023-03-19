package main

import "fmt"

func findMaxMin(data []int) (max int, min int) { //var max, min adalah int
	min = data[0]
	max = data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > max { //proses menentukan angka paling besar
			max = data[i]
		}
		if data[i] < min { //proses menentukan angka paling kecil
			min = data[i]
		}
	}
	return
}

func main() {
	data := []int{4, 7, 1, 10, 30, 6, 8} //value var data akan dikirim ke fun findMaxMin
	max, min := findMaxMin(data)         //value findMaxMin diambil dari var data
	fmt.Println("max:", max, "min:", min)
}
