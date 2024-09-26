package main

import (
	"fmt"
	"lemin/lemin"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatalln("Input file missing!")
	}

	linksMap, start, end, _, _, err := lemin.ReadData(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	AllPossiblePaths, err := lemin.FindAllPaths(linksMap, start, end)
	if err != nil {
		log.Fatalln(err)
	}

	FilteredPaths := lemin.FilterPaths(AllPossiblePaths, start, end)
	fmt.Println(AllPossiblePaths)
	fmt.Print("\n\n\n")
	fmt.Println(FilteredPaths)
}
