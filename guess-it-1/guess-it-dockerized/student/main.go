package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var num int
	arr := []int{}

	for i := 1; i <= 12500; i++ {
		fmt.Fscan(os.Stdin, &num)
		arr = append(arr, num)

		low := Median(arr) - int(float64(MAD(arr))*1.5)
		hight := Median(arr) + int(float64(MAD(arr))*1.5)
		
		fmt.Println(low, hight)
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

func MAD(numbers []int) int {
	med := Median(numbers)

	// Step 2: Calculate the absolute deviations from the median
	absDeviations := make([]int, len(numbers))
	for i, num := range numbers {
		absDeviations[i] = int(math.Abs(float64(num) - float64(med)))
	}

	// Step 3: Calculate the median of the absolute deviations
	return Median(absDeviations)
}