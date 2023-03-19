package main

import "fmt"

func main() {
	//jika deklarasi menggunakan var biasa, cara akses memory menggunakan &
	var name string = "john"
	fmt.Println("name address", &name) //outputnya memory address
	fmt.Println("name value", name)

	//deklarasi menggunakan var pointer
	var namePointer *string = &name         //value memory dari var names
	fmt.Println("namePointer", namePointer) //memory address
	//jika var menggunakan pointer, maka akses valuenya menggunakan *
	fmt.Println("namePointer value", *namePointer) //outputnya value dari var name

	fmt.Println("====================================")
	name = "doe"                         //cara value mengubah var biasa
	*namePointer = "tono"                //cara mengubah value var pointer
	fmt.Println("name address", &name)   //alamat tidak pernah berubah
	fmt.Println("name value", name)      //value name berubah
	fmt.Println("name value", len(name)) //jumlah digit value

	fmt.Println("====================================")
	// zero value menggunakan perkondisian
	number_a := 25
	var number_b *int //number_b tidak punya nilai
	if number_b == nil {
		fmt.Println("number_b is", number_b)
		number_b = &number_a                               //value number_b adalah alamat number_A
		fmt.Println("number_b after init : is", *number_b) //number_b mengakses value number_a
	}
}
