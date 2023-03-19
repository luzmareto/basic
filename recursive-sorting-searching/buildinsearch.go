package main

import (
	"fmt"
	"sort"
)

func main() {
	//build in search tetap akan mencetak nilai walaupun nilai tersebut tidak ada di dalam slice/index
	elements := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	value := 32 //outpunya index ke-6, karena 32 terletak diantar 31 dan 53
	findIndex := sort.SearchInts(elements, value)
	fmt.Println("index", findIndex)
	if value == elements[findIndex] {
		fmt.Println("value ditemukan") //output yang keluar jika value ada didlam index
	} else {
		fmt.Println("value tidak ditemukan") //output yang keluar jika value tidak ada didlam index
	}
}
