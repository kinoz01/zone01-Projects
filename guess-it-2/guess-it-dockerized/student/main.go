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

		if i == 1 {
			fmt.Println(num, num)
			Y = append(Y, num)
			X = append(X, float64(i))
			i++
			continue
		}

		X = append(X, float64(i))
		Y = append(Y, num)

		a, b := LinearRegression(X, Y)
		r := PearsonCorrelation(X, Y)

		fmt.Println((a*float64(i)+b)-(a*float64(i)+b)*(1-r)-4.5, (a*float64(i)+b)+(a*float64(i)+b)*(1-r)+4.5)

		i++
	}
}

func LinearRegression(xValues, yValues []float64) (float64, float64) {
	n := float64(len(xValues))
	var sumX, sumY, sumXY, sumX2 float64

	for i := range xValues {
		sumX += xValues[i]
		sumY += yValues[i]
		sumXY += xValues[i] * yValues[i]
		sumX2 += xValues[i] * xValues[i]
	}

	// Calculate slope (m)
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	// Calculate intercept (b)
	b := (sumY - m*sumX) / n

	return m, b
}

func PearsonCorrelation(xValues, yValues []float64) float64 {
	n := float64(len(xValues))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i := range xValues {
		sumX += xValues[i]
		sumY += yValues[i]
		sumXY += xValues[i] * yValues[i]
		sumX2 += xValues[i] * xValues[i]
		sumY2 += yValues[i] * yValues[i]
	}

	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))

	if denominator == 0 {
		return math.NaN()
	}

	return numerator / denominator
}
