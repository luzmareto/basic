package main

import "fmt"

type Person struct {
	Name   string
	Age    uint
	Height uint
	Weight uint
}

type Mobil struct {
	Jenis      string
	Merk       string
	Transmisi  string
	BahanBakar string
	CC         uint
	Tahun      uint
	Warena     string
}

func main() {
	var person1 Person
	person1.Name = "john"
	person1.Age = 20
	person1.Weight = 70
	person1.Height = 170
	fmt.Println("person1:", person1) //outputnya seluruh struct person

	//short declaration
	person2 := Person{"tono", 100, 12, 14}
	fmt.Println("person2", person2)
	fmt.Println("=============")

	var mobil1 Mobil
	mobil1.Jenis = "SUV"
	mobil1.BahanBakar = "listrik"
	mobil1.CC = 3000
	mobil1.Merk = "esemka"
	mobil1.Tahun = 2024
	mobil1.Warena = "putih"
	fmt.Println("alamat:", &mobil1)           //alamat var mobil
	fmt.Println("mobil1:", mobil1.BahanBakar) //hanya mengambil 1 value dari struct mobil
}
