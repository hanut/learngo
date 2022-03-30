package main

import (
	"fmt"
)

func main() {
	var t uint8
	fmt.Scanln(&t)
	var g, n []uint32 = make([]uint32, t, t), make([]uint32, t, t)
	var inSl [][]uint32 = make([][]uint32, t, t)
	for i := uint8(0); i < t; i++ {
		fmt.Scanln(&n[i])
		fmt.Scanln(&g[i])
		inSl[i] = make([]uint32, g[i], g[i])
		for j := 0; j < int(g[i]-1); i++ {
			_, err := fmt.Scan(&inSl[i][j])
			if err != nil {
				panic(err)
			}
		}

		inSl[i] = QSort(inSl[i], 0, int(g[i]-1))
		sum := uint32(0)
		for _, v := range inSl[i][:n[i]] {
			sum += v
		}
		fmt.Println(sum)
	}
}

func Prt(arr []uint32, l, h int) ([]uint32, int) {
	pv := arr[h]
	i := l
	for j := l; j < h; j++ {
		if arr[j] < pv {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[h] = arr[h], arr[i]
	return arr, i
}

func QSort(arr []uint32, l int, h int) []uint32 {
	if l < h {
		var p int
		arr, p = Prt(arr, l, h)
		arr = QSort(arr, l, p-1)
		arr = QSort(arr, p+1, h)
	}
	return arr
}
