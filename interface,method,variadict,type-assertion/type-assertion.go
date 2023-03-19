package main

import (
	"fmt"
	"strings"
)

func explain(i interface{}) { //interface kosong
	switch i.(type) { //typedata bebas
	case string: //output jika diisi value string
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i) // output jika diisi valuenya int
	default: //output jika mengluarkan type data selain int dan string
		fmt.Println("i stored something else", i)
	}
}

func main() {
	var secret interface{}
	secret = 2
	var number = secret.(int) * 10                      //(int) berfungsi untuk menentukan typedata yang akan dikali 10
	fmt.Println(secret, "multiplied by 10 is:", number) //20

	//menggunakan slice
	secret = []string{"apple", "manggo", "banana"} //SLICE OF STRING
	//gunakan join untuk menggabung slice
	var fruits = strings.Join(secret.([]string), ",") //"," digunakan untuk memberikan koma pada o
	fmt.Println(fruits, "is my favorite fruits")

	var data interface{}
	data = 10
	explain(data) //i stored int 10
	data = "alta"
	explain(data) //i stored string ALTA
	//secret adalah var interface yang sudah dideklarasi diatas,
	explain(secret) // i stored something else [appel manggo banana]
}
