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
            continue
        }
        Y = append(Y, num)
        X = append(X, float64(i))
        avgx := Averge(X)
        avgy := Averge(Y)
        Xi := Center(X, avgx)
        Yi := Center(Y, avgy)
        a := CalculateSlope(Xi, Yi)
        b := avgy - a*avgx
        avgxi := Averge(Xi)
        avgyi := Averge(Yi)
        var1 := Vrariance(Xi, avgxi)
        var2 := Vrariance(Yi, avgyi)
        cov := Cov(Xi, Yi, avgxi, avgyi) / math.Sqrt(var1*var2)
        if cov < 0 {
            cov = -cov
        }
        fmt.Println((a*float64(i)+b)-(a*float64(i)+b)*(1-cov)-4.5, (a*float64(i)+b)+(a*float64(i)+b)*(1-cov)+4.5)
		i++
    }
}

func Averge(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return float64(sum) / float64(len(nums))
}

func Center(Z []float64, avg float64) []float64 {
    res := []float64{}
    for _, z := range Z {
        res = append(res, z-avg)
    }
    return res
}

func CalculateSlope(X []float64, Y []float64) float64 {
    n := 0.0
    N := 0.0
    for i := 0; i < len(X); i = i + 1 {
        n += X[i] * Y[i]
        N += X[i] * X[i]
    }
    // fmt.Println(n, N)
    return n / N
}

func Cov(X, Y []float64, avgx, avgy float64) float64 {
    res := 0.0
    for i := 0; i < len(X); i++ {
        res += (X[i] - avgx) * (Y[i] - avgy)
    }
    return res / float64(len(X))
}

func Vrariance(nums []float64, averg float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += math.Pow((num)-(averg), 2)
    }
    return sum / float64(len(nums))
}