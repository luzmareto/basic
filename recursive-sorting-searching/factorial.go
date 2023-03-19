/*contoh factorial
5! = dibaca 5 faktorial. rumus factorial :
5! = 5 x 4 x 3 x 2 x 1 = 120.
rumus recursive :

5! = 5 x 4!
*/

package main

import "fmt"

// factorial
func factorialLoop(angka int) int { //VALUE ANGKA DIISI DI FUNC MAIN
	//SCRIPT FACTORIAL MENGGUNAKAN FOR
	var result int = 1            //VALUE RESULT ADALAH SATU
	for i := angka; i >= 1; i-- { //PERULANGAN MULAI DARI VALUE ANGKA SAMPAI 1
		result *= i //MEMBUAT OPERASI PERKALIAN DARI VALUE ANGKA SAMPAI ANGKA 1
	}
	return result
}

// recursive adalah func yang memanggil dirinya sendiri
func factorialRecursive(angka int) int {
	// bace case
	if angka == 1 {
		return 1 //jika value angka adalah 1, maka langsung mereturn 1
	} else { //recurrent relation
		return angka * factorialRecursive(angka-1) //value angka * factorialRecursive
	}
}

func main() {
	fmt.Println(factorialLoop(5)) //value yang akan dikirim ke func factorialLoop
	fmt.Println(factorialRecursive(6))
}
