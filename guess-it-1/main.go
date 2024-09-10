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

		rRange(arr)
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

func rRange(arr []int) {
	centr := Median(arr)
	lowerBound := centr - 45
	upperBound := centr + 45
	fmt.Println(lowerBound, upperBound)
}