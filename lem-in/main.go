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

	linksMap, start, end, antsNum, _, err := lemin.ReadData(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	AllPossiblePaths, err := lemin.FindAllPaths(linksMap, start, end)
	if err != nil {
		log.Fatalln(err)
	}

	PathsGroups := lemin.GenerateNonCrossingPathGroups(AllPossiblePaths, start, end)
	//fmt.Println(AllPossiblePaths)
	fmt.Print("\n\n\n")
	for _, paths := range PathsGroups {
		fmt.Println(paths)
	}
	fmt.Print("\n\n\n")

	// Calculate the best path group for the ants
	ok, bestMetrics, bestGroup := lemin.CalculateBestGroup(PathsGroups, antsNum, end)

	// Print the best group and metrics
	fmt.Println("Best Path Group:")
	for _, path := range bestGroup {
		fmt.Println(path)
	}
	fmt.Printf("Lines: %d, Moves: %d\n", bestMetrics.Lines, bestMetrics.Moves)

	for _, s := range ok {
		fmt.Println(s)
	}

}
