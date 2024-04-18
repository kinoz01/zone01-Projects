package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	// "regexp"
)

func replaceText(text *string, old, new string) {
	if text == nil {
		return // if the string pointer is nil, do nothing
	}
	*text = strings.ReplaceAll(*text, old, new)
}


func wordsBeforeFlag(text, flag string) []string {
	var r []string
	words := strings.Fields(text)

	for i, w := range words {
		if w == flag && i > 0{
			r = append(r, words[i-1]) 
		}
	}
	return r
}


func convertFromBaseToBase(s string, a, b int) string {
	
	// Signature: func ParseInt(s string, base int, bitSize int) (i int64, err error)
	num, err := strconv.ParseInt(s, a, 64)
	if err != nil {
		// fmt.Println("Error converting hex to decimal:", err)
		return ""
	}
	return strconv.FormatInt(num, b)
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error usage: <file1.txt> <file2.txt>")
		return
	}

	textBin, err := ioutil.ReadFile(args[0]) 
	if err != nil {
		fmt.Printf("Error reading file: %s\n", args[0])
		return
	}

	fmt.Println(string(textBin))

	text := string(textBin) // Here we have our text as a string

	formatHex(&text)
	formatBin(&text)
	formatLow(&text)
	formatUp(&text)
	formatCap(&text)
	formatLowWithNum(&text)

	fmt.Println(text)
}

func formatHex(text *string) {
	hexSlice := wordsBeforeFlag(*text, "(hex)")  // Here we have a slice of words found before (hex) flag
	for _, aHex := range hexSlice {
		replaceText(text, aHex + " (hex)", convertFromBaseToBase(aHex, 16, 10))
	}
}

func formatBin(text *string) {
	binSlice := wordsBeforeFlag(*text, "(bin)") // Here we have a slice of words found before (bin) flag
	for _, aBin := range binSlice {
		replaceText(text, aBin + " (bin)", convertFromBaseToBase(aBin, 2, 10))
	}
}

func formatLow(text *string) {
	lowSlice := wordsBeforeFlag(*text, "(up)") // Here we have a slice of words found before (up) flag
	for _, aLowWord := range lowSlice {
		replaceText(text, aLowWord + " (up)", strings.ToUpper(aLowWord))
	}
}

func formatUp(text *string) {
	lowSlice := wordsBeforeFlag(*text, "(low)") // Here we have a slice of words found before (up) flag
	for _, anUpWord := range lowSlice {
		replaceText(text, anUpWord + " (low)", strings.ToLower(anUpWord))
	}
}

func formatCap(text *string) {
	words := wordsBeforeFlag(*text, "(cap)") // Here we have a slice of words found before (up) flag
	for _, w := range words {
		replaceText(text, w + " (cap)", strings.Title(w))
	}
}

func formatLowWithNum(text *string) {
	words := lowWithNum(*text, "(low,") // Here we have a slice of words found before (low, flag
	for _, w := range words {
		replaceText(text, w, strings.ToLower(w))
	}
}

func lowWithNum(text, flag string) []string {
	var r []string

	words := strings.Fields(text)
	var numString string

	for i, w := range words {
		if w == flag {
			if i+1 < len(words){
				numString = words[i+1][0:len(words[i+1])-1]
			}
			num, _ := strconv.Atoi(numString)
			for j:=1; j <= num; j++ {
				if i-j < len(words) {
					r = append(r, words[i-j])
				}
			}			
		}
	}
	return r
}