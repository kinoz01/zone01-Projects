package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func readData(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var yValues []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		yValues = append(yValues, val)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(yValues) == 0 {
		return nil, fmt.Errorf("your file don't contain any data")
	}
	return yValues, nil
}

func linearRegression(xValues, yValues []float64) (float64, float64) {
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

func pearsonCorrelation(xValues, yValues []float64) float64 {
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
		return 0
	}

	return numerator / denominator
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path as an argument.")
	}
	filePath := os.Args[1]

	// Read data from file
	yValues, err := readData(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Generate x values (0, 1, 2, ...)
	xValues := make([]float64, len(yValues))
	for i := range xValues {
		xValues[i] = float64(i)
	}

	// Calculate Linear Regression Line
	slope, intercept := linearRegression(xValues, yValues)

	// Calculate Pearson Correlation Coefficient
	pearson := pearsonCorrelation(xValues, yValues)

	// Print results
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearson)
}
