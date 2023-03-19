package main

import "fmt"

type Student struct {
	Nama         string
	JenisKelamin string
	NomorInduk   string
	Jurusan      string
	Telp         []string //slice didalam struct
	Alamat       Address
}

// seluruh atribut address akan dikirim ke atribut alamat pada student struct
type Address struct {
	Jalan string
	nomor int
	Kota  string
}

func main() {
	var student1 Student
	student1.Nama = "adi"
	student1.NomorInduk = "Nis0123"
	student1.Telp = append(student1.Telp, "0123") //append adalah pelengkap dari slice
	student1.Telp = append(student1.Telp, "3210")
	fmt.Println("student1:", student1.Telp[0]) //outputnya telp pertama

	//cara 1
	var alamat1 Address
	alamat1.Jalan = "Jl. Perjuangan"
	alamat1.nomor = 100
	alamat1.Kota = "Jakarta"
	student1.Alamat = alamat1

	//cara 2
	student1.Alamat = Address{
		Jalan: "jl. lurus",
		Kota:  "surabaya",
		nomor: 123,
	}
	fmt.Println("student1:", student1)
	fmt.Println("student1:", student1.Alamat.Jalan) //output jalan
}
