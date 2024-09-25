package main

import (
	"fmt"
)

// Graph representation using adjacency list
type Graph map[string][]string

// Path struct to store the letters in a path, the path itself, and its length
type Path struct {
	letters map[string]bool
	path    []string
	length  int
}

// DFS to find all paths from start to end
func FindPaths(graph Graph, start, end string, visited map[string]bool, path []string, result *[][]string) {
	// Mark the current room as visited
	visited[start] = true
	path = append(path, start)

	// If we reached the end, add the path to the result
	if start == end {
		*result = append(*result, append([]string{}, path...))
	} else {
		// Explore adjacent rooms
		for _, neighbor := range graph[start] {
			if !visited[neighbor] {
				FindPaths(graph, neighbor, end, visited, path, result)
			}
		}
	}

	// Backtrack: unmark the current room and remove it from the current path
	visited[start] = false
	//path = path[:len(path)-1]
}

func main() {
	// Create the graph from the room connections
	graph := Graph{
		"0": {"1", "2"},
		"1": {"4"},
		"2": {"4"},
		"3": {"0"},
		"4": {"3", "5"},
	}

	// Initialize the visited map
	visited := make(map[string]bool)
	for node := range graph {
		visited[node] = false
	}

	// Result to store all paths
	var result [][]string

	// Find all paths from "start" to "end"
	FindPaths(graph, "0", "5", visited, []string{}, &result)

	// Print all paths
	fmt.Println("Possible paths from start to end:")
	for _, path := range result {
		fmt.Println(path)
	}

	// Call the filtering function
	filteredPaths := FilterPaths(result)
	// Print all paths
	fmt.Println("Final possible paths from start to end:")
	for _, path := range filteredPaths {
		fmt.Println(path)
	}
}

func FilterPaths(paths [][]string) [][]string {
	N := len(paths)
	pathInfos := make([]Path, N)

	// Build pathInfos with letters, path, and length
	for i, path := range paths {
		letters := make(map[string]bool)
		for _, letter := range path {
			if letter != "start" && letter != "end" {
				letters[letter] = true
			}
		}
		pathInfos[i] = Path{
			letters: letters,
			path:    path,
			length:  len(path),
		}
	}

	maxPaths := 0
	minTotalLength := 0
	var bestSubset []int

	// Generate all subsets of paths
	for subset := 1; subset < (1 << N); subset++ {
		selectedPaths := []int{}
		totalLength := 0
		lettersUsed := make(map[string]bool)
		valid := true

		for i := 0; i < N; i++ {
			if (subset & (1 << i)) != 0 {
				// Include path i
				selectedPaths = append(selectedPaths, i)
				totalLength += pathInfos[i].length

				// Check for overlapping letters
				for letter := range pathInfos[i].letters {
					if lettersUsed[letter] {
						valid = false
						break
					}
					lettersUsed[letter] = true
				}
				if !valid {
					break
				}
			}
		}

		if valid {
			numPaths := len(selectedPaths)
			if numPaths > maxPaths || (numPaths == maxPaths && totalLength < minTotalLength) {
				maxPaths = numPaths
				minTotalLength = totalLength
				bestSubset = append([]int(nil), selectedPaths...)
			}
		}
	}

	// Build the result using the best subset
	result := [][]string{}
	for _, idx := range bestSubset {
		result = append(result, paths[idx])
	}

	return result
}
