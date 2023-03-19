package main

import "fmt"

/*

menjumlah angka yang ada didalam arrya
misalkan ada 6 data dalam aray [1,2,3,4,10,111]
kita harus menjumlahkan setiap data tesebut dengan rumus
1 + 2 + 3 + 4 + 10 + 111 = 131

*/

func simpleArraySum(ar []int32) int32 { //value func simpleArray Sum diisi pada func main
	var temp int32
	for i := 0; i < len(ar); i++ {
		temp += ar[i] //setiap data akan ditambah data lainya
	}
	return temp
}

func main() {
	data := []int32{1, 2, 3, 4, 10, 111} //value data akan dimasukan ke fun simpleArraySum
	fmt.Println(simpleArraySum(data))
}
