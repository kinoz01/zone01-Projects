package lemin

import (
	"fmt"
	"sort"
)

// PathMetrics holds the metrics for the paths.
type PathMetrics struct {
	Lines int
	Moves int
}

// PathInfo holds the information about each path.
type PathInfo struct {
	Path   []string
	Length int
}

// CalculateBestGroup evaluates the best path group and assigns ants optimally.
func CalculateBestGroup(groups [][][]string, numAnts int, start, end string) ([]string, PathMetrics, [][]string) {
	var bestSteps []string
	bestMetrics := PathMetrics{Lines: int(^uint(0) >> 1)} // Initialize to max
	var bestGroup [][]string

	for _, group := range groups {
		paths := group
		steps, lines, moves := GenerateAntSteps(paths, numAnts, start, end)

		// Update best metrics if this group is better
		if lines < bestMetrics.Lines || (lines == bestMetrics.Lines && moves < bestMetrics.Moves) {
			bestMetrics = PathMetrics{Lines: lines, Moves: moves}
			bestSteps = steps
			bestGroup = paths
		}
	}
	return bestSteps, bestMetrics, bestGroup
}

// GenerateAntSteps assigns ants optimally to paths and generates the steps.
func GenerateAntSteps(paths [][]string, numAnts int, start, end string) ([]string, int, int) {
	// Prepare paths info
	var pathInfos []PathInfo
	for _, path := range paths {
		pathInfos = append(pathInfos, PathInfo{
			Path:   path,
			Length: len(path),
		})
	}

	// Sort paths by length (ascending)
	sort.Slice(pathInfos, func(i, j int) bool {
		return pathInfos[i].Length < pathInfos[j].Length
	})

	// Calculate the number of ants per path
	numPaths := len(pathInfos)
	antsPerPath := make([]int, numPaths)

	// Assign ants to paths to minimize the total time (lines)
	totalAnts := numAnts
	for {
		minTime := pathInfos[0].Length + antsPerPath[0] - 1
		minIndex := 0
		for i := 1; i < numPaths; i++ {
			time := pathInfos[i].Length + antsPerPath[i] - 1
			if time < minTime {
				minTime = time
				minIndex = i
			}
		}
		if totalAnts == 0 {
			break
		}
		antsPerPath[minIndex]++
		totalAnts--
	}

	// Calculate total lines
	lines := 0
	for i := 0; i < numPaths; i++ {
		time := pathInfos[i].Length + antsPerPath[i] - 1
		if time > lines {
			lines = time
		}
	}

	// Simulate ant movements
	type Ant struct {
		ID     int
		Path   []string
		Pos    int
		PathID int
	}

	var steps []string
	ants := make([]Ant, 0, numAnts)
	antID := 1

	// Initialize ant queues per path
	for i, count := range antsPerPath {
		for j := 0; j < count; j++ {
			ant := Ant{
				ID:     antID,
				Path:   pathInfos[i].Path,
				Pos:    -j - 1, // Start position before the first room after 'start'
				PathID: i,
			}
			antID++
			ants = append(ants, ant)
		}
	}

	totalMoves := 0
	for step := 0; step < lines; step++ {
		line := ""
		occupied := make(map[string]bool)
		for i := range ants {
			ant := &ants[i]
			if ant.Pos >= len(ant.Path)-1 {
				continue // Ant has finished its path
			}
			if ant.Pos+1 >= 0 {
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
				totalMoves++
			} else {
				ant.Pos++ // Ant is waiting to start
			}
		}
		if line != "" {
			steps = append(steps, line)
		}
	}

	return steps, lines, totalMoves
}
