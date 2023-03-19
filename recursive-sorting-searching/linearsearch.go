package main

import "fmt"

func linierSearch(elements []int, x int) int { //index slice disamping diisi pada func main
	var count int
	for i := 0; i < len(elements); i++ { //i adalah 0 karna index juga dimulai dari 0
		count++                     //index pertama menjadi satu
		fmt.Println("count", count) //count membaca index pertama dihitung menjadi 1
		if elements[i] == x {
			return i //jika value index sudah ditemukan, maka akan dikembalikan ke i
		}
	}
	return -1 //jika value index tidak ditemukan, maka mencetak -1
}

func main() {
	var data = []int{4, 1, 6, 8, 9, 10, 30} //data adalah kumpulan int index
	result := linierSearch(data, 6)         //dibaca menggunakan index
	if result != -1 {
		fmt.Println("data ditemukan di index ke", result)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
