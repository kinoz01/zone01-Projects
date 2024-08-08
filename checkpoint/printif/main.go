package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(s string) string {
	if len([]rune(s)) > 3 || len([]rune(s)) == 0 {
		return "G\n"
	} else {
		return "Invalid Output\n"
	}
}
