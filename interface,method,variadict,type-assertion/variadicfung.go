package main

import "fmt"

func Cetak(angka ...int) { //angka adalah parameter yang dapat diisi banyak int
	for _, v := range angka {
		fmt.Println(v)
	}

}

func main() {
	a, b, c, d := 10, 20, 30, 40
	Cetak(a, b, c, d) //seluruh value dibaca sebagai 1 parameter yaitu angka
	var data = []int{1, 2, 3, 4}
	fmt.Println(data[0], data)

}
