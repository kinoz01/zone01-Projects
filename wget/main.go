package main

import (
	"os"
	"wget/wget"
)

func main() {

	if len(os.Args) == 1 {
		wget.PrintMissingURL()
	}
	wget.ParseArgs(os.Args[1:])
	wget.Run()
}
