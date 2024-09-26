package lemin

import (
	"errors"
)

// Function to initiate the search
func FindAllPaths(graph map[string][]string, start, end string) ([][]string, error) {
	visited := make(map[string]bool)
	var result [][]string
	FindPaths(graph, start, end, visited, []string{}, &result)
	
	if len(result) == 0 {
		return nil, errors.New("no valid paths found")
	}
	return result, nil
}

// FindPaths finds all paths from start to end in a directed graph.
func FindPaths(graph map[string][]string, start, end string, visited map[string]bool, path []string, result *[][]string) {

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
}
