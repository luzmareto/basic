package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

// method
func (r Rect) HitungLuas() float64 {
	return r.width * r.height //returnya adalah rumus luas persegi
}

// method
func (c Circle) HitungLuas() float64 {
	return math.Pi * c.radius * c.radius //returnya adalah rumus luas lingkaran
}

func main() {
	//mendeklarasi value baru untuk memberi value method di atas
	rect := Rect{5.0, 4.0}                         //var rect typedatanya adalah Rect dengan value width = 5.0 dan height = 4.0
	cir := Circle{5.0}                             //var cir typedatanya adalah Circle dengan value raius = 5.0
	fmt.Println("luas rect =", rect.HitungLuas())  //method Hitungluas akan mengoperasikan value yang ada pada var rect
	fmt.Println("luas circle =", cir.HitungLuas()) //method Hitungluas akan mengoperasikan value yang ada pada var cir
}
