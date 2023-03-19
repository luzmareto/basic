package main

import "fmt"

//anon func adalah func yang tidak memilik nama

func main() {
	value := func() {
		fmt.Println("welcome to geeks")
	}
	value()
	func() {

		fmt.Println("welcome to geeks guys")
	}()
}
