package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func (P Person) GetName() string { //p adalah var dengan typedata Person
	P.name = "andi"             //p.name adalah var untuk memberi value atribut name pada typedata Person
	return P.name + " amazing!" //value name ditambah "amazing"
}

func (P *Person) IncreaseAge() {
	P.age = P.age + 1 //p.age adlah var untuk memberi value atribut name pada typedata Person
	//valuenya ditambah 1
}

func main() {
	//mendeklarasi untuk mengisi value method diatas
	PersonA := Person{"john", 50}  //memberi value
	fmt.Println(PersonA)           //mencetak value diatas
	fmt.Println(PersonA.GetName()) //var GetName akan menampilkan value dari p.name
	fmt.Println(PersonA.name)      //outpunya diambil dari var PersonA yaitu "john"

	PersonB := Person{"doe", 40}   //memberi value
	fmt.Println(PersonB)           //mencetak value diatas
	fmt.Println(PersonB.GetName()) //karena ada var GetName, maka outputnya diambil dari p.name

	fmt.Println("sebelum:", PersonA.age) //outputnya 50, diambil dari var PersonA
	PersonA.IncreaseAge()                //50 + 1
	fmt.Println(PersonA.age)             //51
}
