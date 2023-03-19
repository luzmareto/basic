package main

import "fmt"

type Employee struct {
	FirstName, LastName string
	Age                 uint
}

// function
func fullName(FirstName string, lastName string) (fullName string) {
	fullName = FirstName + " " + lastName
	return //returnya fullname adalah gabungan value dari firstname dan lastname
}

// method
func (e Employee) fullNameMethod() string {
	//versi mudah //returnya fullname adalah gabungan value dari firstname dan lastname
	return e.FirstName + " " + e.LastName
}

func (e Employee) IncreaseAge(angka int) uint { //e adalah variabel baru yang isinya atribut dari employee
	result := e.Age + uint(angka) //20 adalah age, 10 adalah angka dari var temo
	return result                 //mengembalikan value yang ditambah value angka
}

func main() {
	//mendeklarasi sekaligus memberi value firstname, lastname dan age
	emp1 := Employee{
		FirstName: "ros",
		LastName:  "geler",
		Age:       20,
	}
	var temp = 10 //temp dipakai untuk value angka

	fmt.Println(fullName(emp1.FirstName, emp1.LastName))
	fmt.Println(emp1.fullNameMethod())  //fullname versi method
	fmt.Println(emp1.IncreaseAge(temp)) // 20 + 10
}
