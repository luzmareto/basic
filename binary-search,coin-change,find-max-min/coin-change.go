package main

import (
	"fmt"
	"sort"
)

/*
coin-change adalah menukar value dengan pecahan angka yang ada didalam slice
congtoh analogi:
kita punya koin(value) sebanyak 24
sementara di dalam box(slice) ada pecahan koin 10, 25, 5, 1
berarti kita dapat menukar 24 koin menjadi 10 10 1 1 1 1 totalnya 42
kita dapat menukar dari urutan terbesar ke terkecil atau sebaliknya
*/

func coinChange(money int, coinvalue []int) []int {
	sort.SliceStable(coinvalue, func(i, j int) bool { //membuat urutan value dalam sl
		return coinvalue[i] > coinvalue[j] //membuat urutan dari angka terbesar
	})
	result := []int{}
	for _, coin := range coinvalue {
		if money >= coin {
			for money >= coin {
				result = append(result, coin)
				money = money - coin
			}
		} else {
			continue
		}
	}

	return result
}

func main() {
	coins := []int{10, 25, 5, 1}
	fmt.Println(coinChange(24, coins))
}
