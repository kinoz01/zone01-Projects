package main

import (
	"fmt"
	"os"
)

const (
	dataSize = 2048
	maxOps   = 4096
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: brainfuck <code>")
		os.Exit(1)
	}

	code := os.Args[1]
	if len(code) > maxOps {
		fmt.Println("Code too long")
		os.Exit(1)
	}

	data := make([]byte, dataSize)
	ptr := 0
	ip := 0

	for ip < len(code) && ip >= 0 {
		switch code[ip] {
		case '>':
			ptr++
			if ptr >= dataSize {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = dataSize - 1
			}
		case '+':
			data[ptr]++
		case '-':
			data[ptr]--
		case '.':
			fmt.Printf("%c", data[ptr])
		case '[':
			if data[ptr] == 0 {
				level := 1
				for ip++; ip < len(code) && level > 0; ip++ {
					switch code[ip] {
					case '[':
						level++
					case ']':
						level--
					}
				}
				if level != 0 {
					fmt.Println("Unmatched '['")
					os.Exit(1)
				}
			}
		case ']':
			if data[ptr] != 0 {
				level := 1
				for ip--; ip >= 0 && level > 0; ip-- {
					switch code[ip] {
					case '[':
						level--
					case ']':
						level++
					}
				}
				if level != 0 {
					fmt.Println("Unmatched ']'")
					os.Exit(1)
				}
				ip++ // Skip the ']' again
			}
		}
		ip++
	}
}
