package main

import (
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 || len(os.Args[1]) == 0 || len(os.Args[2]) == 0 {
		return
	}

	expressions := CheckEXpression(os.Args[1])
	if expressions == nil {
		return
	}
	words := Fields(os.Args[2], ' ')

	num := 0

	for _, word := range words {
		for _, exp := range expressions {
			if Contains(word, exp) {
				num++
				n := Itoa(num)
				os.Stdout.WriteString(n + ": " + word + "\n")
			}
		}
	}
}

func CheckEXpression(s string) []string {
	if s[0] != '(' || s[len([]rune(s))-1] != ')' {
		return nil
	}
	str := s[1 : len(s)-1]
	for _, char := range str {
		if char == '(' || char == ')' {
			return nil
		}
	}
	return Fields(str, '|')
}

func Fields(s string, c rune) []string {
	var res []string
	var chars []rune
	var result []string

	for _, char := range s {
		if char == c {
			res = append(res, string(chars))
			chars = []rune{}
		} else {
			chars = append(chars, char)
		}
	}
	res = append(res, string(chars))
	for _, str := range res {
		if str == string(c) || len(str) == 0 || str == "" {
			continue
		}
		result = append(result, str)
	}
	return result
}

func Itoa(n int) (s string) {
	for n > 0 {
		s = string(rune(n%10)+'0') + s
		n /= 10
	}
	return s
}

func Contains(s, subs string) bool {
	for i := 0; i <= len(s)-len(subs); i++ {
		if s[i:i+len(subs)] == subs {
			return true
		}
	}
	return false
}
