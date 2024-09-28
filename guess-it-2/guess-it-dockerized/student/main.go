package main

import (
	"fmt"
	"math"
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
		a := LinearRegression(X, Y)

		avgx2 := Median(X)
        avgy2 := Median(Y)
		b2 := avgy2 - a*avgx2

		// Recompute regression with filtered data
		r := PearsonCorrelation(X, Y)

		// Calculate standard deviation of Y
		s_y := StandardDeviation(Y)

		// Calculate standard error of estimate (SEE)
		SEE := s_y * math.Sqrt(1 - r*r)

		// Set k value to adjust confidence level (e.g., 0.2 for ~12% confidence)
		k := 0.2

		// Predicted Y value
		yPred := a*float64(i) + b2

		// Compute prediction interval
		// lower := yPred - k*SEE
		//upper := yPred + k*SEE

		fmt.Println(yPred - k*SEE, yPred + k*SEE)

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

// PearsonCorrelation calculates the Pearson correlation coefficient
func PearsonCorrelation(X, Y []float64) float64 {
	n := float64(len(X))
	var sumX, sumY, sumXY, sumX2, sumY2 float64
	for i := 0; i < len(X); i++ {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		sumX2 += X[i] * X[i]
		sumY2 += Y[i] * Y[i]
	}
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// StandardDeviation calculates the standard deviation of a slice
func StandardDeviation(data []float64) float64 {
	n := float64(len(data))
	if n == 0 {
		return 0
	}
	var sum, mean, variance float64
	for _, v := range data {
		sum += v
	}
	mean = sum / n
	for _, v := range data {
		variance += (v - mean) * (v - mean)
	}
	return math.Sqrt(variance / n)
}

func Averge(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return float64(sum) / float64(len(nums))
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
