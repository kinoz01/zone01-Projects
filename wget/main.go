package main

import (
	"os"
	"sync"
	"wget/wget"
)

func main() {

	if len(os.Args) == 1 {
		wget.PrintMissingURL()
	}
	wget.ParseArgs(os.Args[1:])
	if !wget.Mirror {
		wget.Run()
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		wget.MirrorWebsite(&wg)
		wg.Wait()
	}
}
