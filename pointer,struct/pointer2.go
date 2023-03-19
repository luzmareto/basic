package main

import "fmt"

func ubahAngka(bil int) int {
	fmt.Println("bil:", bil) //90
	bil = 100                //value bil diubah jadi 100
	return bil
}

func ubahAngkaPointer(bil *int) int {
	fmt.Println("bil:", *bil) //90
	fmt.Println("alamat bil:", *bil)
	*bil = 200 //value bil diubah jadi 200
	return *bil

}

func main() {
	var nilai = 90
	fmt.Println("ubahAngka:", ubahAngka(nilai)) //value bil adalah 90 lalu diubah jadi 100
	fmt.Println("nilai", nilai)
	fmt.Println("========================")
	//mengubah value pada var pointer
	fmt.Println("ubahAngkaPointer:", ubahAngkaPointer(&nilai))
	fmt.Println("nilai", nilai)
	fmt.Println("alamat nilai", &nilai)
}
