package main

import (
	"fmt"
	"math"
	"os"
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

		if i == 1 {
			fmt.Println(num, num)
			i++
			continue
		}

		// Perform linear regression on current data
		a, b := LinearRegression(X, Y)

		// Calculate residuals
		residuals := make([]float64, len(Y))
		for j := 0; j < len(Y); j++ {
			yPred := a*X[j] + b
			residuals[j] = Y[j] - yPred
		}

		// Calculate standard deviation of residuals
		residualStd := StandardDeviation(residuals)

		// Identify outliers (residuals greater than 2 times the standard deviation)
		filteredX := []float64{}
		filteredY := []float64{}
		for j := 0; j < len(Y); j++ {
			if math.Abs(residuals[j]) <= 1.4*residualStd {
				filteredX = append(filteredX, X[j])
				filteredY = append(filteredY, Y[j])
			}
		}

		// Ensure we have enough data points after filtering
		if len(filteredX) < 2 {
			i++
			continue
		}

		// Recompute regression with filtered data
		a, b = LinearRegression(filteredX, filteredY)
		r := PearsonCorrelation(filteredX, filteredY)

		// Calculate standard deviation of Y
		s_y := StandardDeviation(filteredY)

		// Calculate standard error of estimate (SEE)
		SEE := s_y * math.Sqrt(1 - r*r)

		// Set k value to adjust confidence level (e.g., 0.2 for ~12% confidence)
		k := 0.21

		// Predicted Y value
		yPred := a*float64(i) + b

		// Compute prediction interval
		lower := yPred - k*SEE
		upper := yPred + k*SEE

		fmt.Println(lower, upper)

		i++
	}
}

// LinearRegression calculates the slope (a) and intercept (b)
func LinearRegression(X, Y []float64) (a, b float64) {
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
	b = (sumY - a*sumX) / n
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
