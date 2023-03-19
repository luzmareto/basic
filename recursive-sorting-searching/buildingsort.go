package main

import (
	"fmt"
	"sort"
)

func main() {
	//sorting dari karakter awal terkecil
	strs := []string{"cb", "ca", "bc"}
	sort.Strings(strs) //sort.Strings untuk mensortir value slice di var strs
	fmt.Println("strings:", strs)

	// sorting dari angka terkecil
	ints := []int{7, 2, 4}
	sort.Ints(ints) //sort.Ints untuk mensorti value slice int pada var ints
	fmt.Println("ints:	", ints)

	//sort.intAreSorted untuk mengecek apakah var sudah disortir atau belum dengan output bool
	s := sort.IntsAreSorted(ints) //karena var ints sudah tersortir, maka outputnya tru
	fmt.Println("sorted :", s)

	//mensortir dari angka terbesar
	gede := []int{12, 5, 1, 2, 51}
	sort.Sort(sort.Reverse(sort.IntSlice(gede)))
	fmt.Println(gede)

	/*
		< adalah ascending(dari kecil ke besar)
		> adalah descending(dari besar ke kecil)
	*/
	keys := []string{"sd", "as"}
	//mensortir tanpa aturan int/string
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j] // bisa memakai ">" atau "<" tergantung kebutuhan
	})
	fmt.Println(keys)
}
