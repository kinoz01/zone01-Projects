package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var num float64
	Y := []float64{}
	X := []float64{}
	i := 1

	for {
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			os.Exit(0)
		}

		X = append(X, float64(i))
		Y = append(Y, num)

		// Ensure we have enough data points after filtering
		if len(X) < 2 {
			i++
			continue
		}
		
		a := LinearRegression(X, Y) // approximating the data slope (didn't use LAD because server crash due to time complexity)
		
		// using Median Regression (Least Absolute Deviations or LAD) for the intercept
		X_median := Median(X)
		Y_median := Median(Y)

		b := Y_median - a*X_median  

		y_approx := a*float64(i) + b

		fmt.Println(y_approx-20, y_approx+20)

		i++
	}
}

// LinearRegression calculates the slope (a) and intercept (b)
func LinearRegression(X, Y []float64) (a float64) {
	n := float64(len(X))
	var sumX, sumY, sumXY, sumX2 float64
	for i := 0; i < len(X); i++ {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		sumX2 += X[i] * X[i]
	}
	denominator := n*sumX2 - sumX*sumX
	if denominator == 0 {
		a = 0
	} else {
		a = (n*sumXY - sumX*sumY) / denominator
	}
	return
}

// Calculates the median of an array of integers.
func Median(arr []float64) float64 {
	le := len(arr)
	sort.Float64s(arr)

	if le%2 == 1 {
		return (arr[(le-1)/2])
	} else {
		return (arr[le/2] + arr[le/2-1]) / 2
	}
}

// Function to calculate the median slopes
// ROBUST to outliers but big time complexity (we are going to use linear regression for the slope)
func MedianSlope(X, Y []float64) float64 {
	if len(X) != len(Y) {
		panic("X and Y must have the same length")
	}

	var slopes []float64

	// Calculate the slope for all pairs (i, j) where i != j
	for i := 0; i < len(X); i++ {
		for j := i + 1; j < len(X); j++ {
			// Avoid division by zero for vertical lines (X_j == X_i)
			if X[j] != X[i] {
				slope := (Y[j] - Y[i]) / (X[j] - X[i])
				slopes = append(slopes, slope)
			}
		}
	}

	// Return the median of the slopes
	return Median(slopes)
}
