package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"
	for i := 0; i < len(arr); i++ {
		z01.PrintRune(rune(hex[arr[i]>>4]))
		z01.PrintRune(rune(hex[arr[i]&15]))
		z01.PrintRune(' ')
		if (i-3)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		}
	}
	for _, byt := range arr {
		if byt < 32 || byt > 128 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(byt))
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}