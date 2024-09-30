package lemin

import (
	"fmt"
	"sync"
)

// PathMetrics holds the metrics for the paths.
type PathMetrics struct {
	Lines int
	Moves int
}

// Generate all possible ant assignments for a given number of ants and paths.
func GenerateAntAssignments(numAnts, numPaths int) [][]int {
	var result [][]int
	var helper func(remainingAnts, index int, current []int)
	helper = func(remainingAnts, index int, current []int) {
		if index == numPaths-1 {
			// Last path gets all remaining ants
			current = append(current, remainingAnts)
			result = append(result, append([]int(nil), current...))
			return
		}
		for i := 0; i <= remainingAnts; i++ {
			helper(remainingAnts-i, index+1, append(current, i))
		}
	}
	helper(numAnts, 0, []int{})
	return result
}

// GenerateAntSteps generates the step sequences for a given group of paths and ant assignment.
func GenerateAntSteps(paths [][]string, antAssignment []int, end string) ([]string, int, int) {
	type Ant struct {
		ID   int
		Path []string
		Pos  int
	}

	// Assign ants to paths based on the antAssignment
	ants := make([]Ant, 0)
	antID := 1
	for i, antCount := range antAssignment {
		for j := 0; j < antCount; j++ {
			ants = append(ants, Ant{
				ID:   antID,
				Path: paths[i],
				Pos:  0,
			})
			antID++
		}
	}

	numAnts := len(ants)
	var steps []string
	totalMoves := 0
	antsFinished := 0

	for antsFinished < numAnts {
		line := ""
		occupied := make(map[string]bool) // Reset occupied rooms each time step
		for i := range ants {
			ant := &ants[i]
			if ant.Pos >= len(ant.Path)-1 {
				continue // Ant has finished its path
			}

			nextRoom := ant.Path[ant.Pos+1]

			// Check if the next room is occupied, except for 'end'
			if nextRoom != end && occupied[nextRoom] {
				continue // Can't move, room is occupied
			}

			// Move the ant
			ant.Pos++
			if nextRoom != end {
				occupied[nextRoom] = true
			}
			line += fmt.Sprintf("l%d-%s ", ant.ID, nextRoom)
			if nextRoom == end {
				antsFinished++
			}
			totalMoves++
		}
		if line != "" {
			steps = append(steps, line)
		}
	}
	lines := len(steps)
	return steps, lines, totalMoves
}

// CalculateBestGroup evaluates all path groups and returns the best one.
func CalculateBestGroup(groups [][][]string, numAnts int, end string) ([]string, PathMetrics, [][]string) {
	var bestSteps []string
	bestMetrics := PathMetrics{Lines: int(^uint(0) >> 1), Moves: int(^uint(0) >> 1)} // Initialize to max
	var bestGroup [][]string

	var mu sync.Mutex      // Mutex to protect shared variables
	var wg sync.WaitGroup  // WaitGroup to wait for all goroutines to finish

	for _, group := range groups {
		wg.Add(1)
		go func(paths [][]string) {
			defer wg.Done()
			antAssignments := GenerateAntAssignments(numAnts, len(paths))
			var innerWg sync.WaitGroup
			for _, assignment := range antAssignments {
				// Skip assignments where total ants assigned is not equal to numAnts
				assignedAnts := 0
				for _, count := range assignment {
					assignedAnts += count
				}
				if assignedAnts != numAnts {
					continue
				}

				innerWg.Add(1)
				go func(assignment []int) {
					defer innerWg.Done()
					steps, lines, moves := GenerateAntSteps(paths, assignment, end)

					// Update best metrics if this group is better
					mu.Lock()
					if lines < bestMetrics.Lines || (lines == bestMetrics.Lines && moves < bestMetrics.Moves) {
						bestMetrics = PathMetrics{Lines: lines, Moves: moves}
						bestSteps = steps
						bestGroup = paths
					}
					mu.Unlock()
				}(append([]int(nil), assignment...)) // Pass a copy of assignment
			}
			innerWg.Wait()
		}(group)
	}
	wg.Wait()
	return bestSteps, bestMetrics, bestGroup
}
