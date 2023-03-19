package main

import "fmt"

//binary-search lebih efesien dalam menghitung perulangan pencarian data

func binarySearch(data []int, x int) int {
	var kiri = 0
	var kanan = len(data) - 1
	var count = 0       //var count untuk informasi perulangan
	for kiri <= kanan { //agar perulangan dari kanan dan kiri tidak bertabrakan
		var tengah = (kiri + kanan) / 2
		count++ //var count++ untuk informasi perulangan
		fmt.Println("count", count)
		if x < data[tengah] {
			kanan = tengah - 1
		} else if x > data[tengah] {
			kiri = tengah + 1
		} else if data[tengah] == x {
			return tengah
		}
	}
	return -1
}

func main() {
	data := []int{1, 4, 7, 9, 10, 14, 17, 20, 24}
	fmt.Println(binarySearch(data, 10))
}
