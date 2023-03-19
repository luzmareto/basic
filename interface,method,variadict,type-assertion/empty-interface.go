package main

import "fmt"

func describe(i interface{}) { //return interfacenya kosong, artinya bisa menggunakan semua typedata
	fmt.Printf("(%v, %T)\n", i, i)
	// %v untuk value. %T untuk typedata. jumlah i harus sesuai jumlah format
}

func main() {
	var data interface{} //typedata data adalah interface describe
	data = "alta"
	describe(data) //mencetak value data dengan interface describe
	data = 10
	describe(data)
	data = true
	describe(data)

	//interface kosong + map dengan tanpa atribut typedata
	var data1 = map[string]interface{}{}
	data1["nama"] = "john"
	data1["umur"] = 20
	data1["isAktif"] = true
	fmt.Println(data1)
}
