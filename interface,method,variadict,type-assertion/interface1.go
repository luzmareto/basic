package main

import (
	"fmt"
)

//pada interfaces, seluruh atribut harus memiliki function terpisah

type Calculate interface {
	large() int
	keliling() int
}

type Square struct { //value square/side diisi pada function main
	side int
}

type Lingkaran struct {
	jari int
}

func (s Square) large() int { // s mengambil fungsi square dan side
	return s.side * s.side //s mengambil value dari side lalu melakukan operasi
}

func (s Square) keliling() int {
	return 4 * s.side //4 x value yang diberikan kepada atribut side pada typedata Square
}

func (l Lingkaran) large() int {
	return 3 * l.jari * l.jari
}

func (l Lingkaran) keliling() int {
	return 4 * l.jari
}

func main() {
	var dimResult Calculate                  //typedata dimResult adalah Calculate
	dimResult = Square{10}                   //memberi value 10 pada atribut Side yang ada di typedata Square
	var hasil = dimResult.large()            //outputnya dimResult melakukan operasi yang ada pada return Square
	fmt.Println("hasil :", hasil)            //100
	var hasilKeliling = dimResult.keliling() //dimResult mengambil rumus yang ada pada function keliling
	fmt.Println("hasil:", hasilKeliling)     //40 didapat karna value square adalah 10

	var dimResultLingkaran Calculate
	dimResultLingkaran = Lingkaran{10}
	var hasilLingkaran = dimResultLingkaran.large()
	fmt.Println("hasil lingkaran:", hasilLingkaran)
}
