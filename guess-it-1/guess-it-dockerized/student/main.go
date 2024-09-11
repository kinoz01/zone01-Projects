package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var num int
	arr := []int{}

	for i := 1; i <= 12500; i++ {
		fmt.Fscan(os.Stdin, &num)
		arr = append(arr, num)

		low := Median(arr) - 45
		hight := Median(arr) + 45
		if i < 2000 {
			fmt.Println(num, num)
			continue
		}
		if i >= 2000 {
			fmt.Println(low, hight)
		}
	}
}

func Median(arr []int) int {
	le := len(arr)
	sort.Ints(arr)

	if le%2 == 1 {
		return (arr[(le-1)/2])
	} else {
		return (arr[le/2] + arr[le/2-1]) / 2
	}
}
