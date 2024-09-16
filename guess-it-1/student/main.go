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

	for {
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			os.Exit(0)
		}
		arr = append(arr, num)

		low := Median(arr) - int(float64(MAD(arr))*1.5)
		hight := Median(arr) + int(float64(MAD(arr))*1.5)

		fmt.Println(low, hight)
	}
}

// Calculates the median of an array of integers.
func Median(arr []int) int {
	le := len(arr)
	sort.Ints(arr)

	if le%2 == 1 {
		return (arr[(le-1)/2])
	} else {
		return (arr[le/2] + arr[le/2-1]) / 2
	}
}

// Calculates the Median Absolute Deviation (MAD) of an array of integers.
// MAD is a robust measure of statistical dispersion.
func MAD(numbers []int) int {
	med := Median(numbers)

	// Calculate the absolute deviations from the median
	absDeviations := make([]int, len(numbers))
	for i, num := range numbers {
		absDeviations[i] = int(math.Abs(float64(num) - float64(med)))
	}

	// Calculate the median of the absolute deviations
	return Median(absDeviations)
}
