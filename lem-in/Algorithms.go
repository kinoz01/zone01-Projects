package main


// Function to filter paths based on the given rules
func FilterPathsFDS(paths [][]string) [][]string {
	// Step 1: Map letters to indices
	letterToIndex := make(map[string]int)
	letterIndex := 0
	for _, path := range paths {
		for _, letter := range path {
			if letter != "start" && letter != "end" {
				if _, exists := letterToIndex[letter]; !exists {
					letterToIndex[letter] = letterIndex
					letterIndex++
				}
			}
		}
	}

	// Number of unique letters
	M := letterIndex

	// Step 2: Precompute path letter bitsets
	type PathInfo struct {
		index         int
		path          []string
		lettersBitset []uint64
		length        int
	}

	pathsInfo := make([]PathInfo, len(paths))
	bitsetSize := (M + 63) / 64

	for i, path := range paths {
		lettersBitset := make([]uint64, bitsetSize)
		length := len(path)
		for _, letter := range path {
			if letter != "start" && letter != "end" {
				idx := letterToIndex[letter]
				quotient := idx / 64
				remainder := idx % 64
				lettersBitset[quotient] |= 1 << remainder
			}
		}
		pathsInfo[i] = PathInfo{
			index:         i,
			path:          path,
			lettersBitset: lettersBitset,
			length:        length,
		}
	}

	// Initialize best selection variables
	var bestSelection []int
	maxPaths := 0
	minTotalLength := 0

	// Step 3: Implement DFS
	var dfs func(index int, selectedPaths []int, usedLetters []uint64)
	dfs = func(index int, selectedPaths []int, usedLetters []uint64) {
		if index == len(pathsInfo) {
			currentPaths := len(selectedPaths)
			if currentPaths > maxPaths {
				maxPaths = currentPaths
				bestSelection = append([]int(nil), selectedPaths...)
				// Calculate total length
				minTotalLength = 0
				for _, idx := range selectedPaths {
					minTotalLength += pathsInfo[idx].length
				}
			} else if currentPaths == maxPaths {
				// Compare total length
				currentTotalLength := 0
				for _, idx := range selectedPaths {
					currentTotalLength += pathsInfo[idx].length
				}
				if currentTotalLength < minTotalLength {
					minTotalLength = currentTotalLength
					bestSelection = append([]int(nil), selectedPaths...)
				}
			}
			return
		}

		pathInfo := pathsInfo[index]

		// Check if we can include this path
		conflict := false
		for i := 0; i < bitsetSize; i++ {
			if usedLetters[i]&pathInfo.lettersBitset[i] != 0 {
				conflict = true
				break
			}
		}

		// Option 1: Exclude the path
		dfs(index+1, selectedPaths, usedLetters)

		// Option 2: Include the path if no conflict
		if !conflict {
			// Update usedLetters
			newUsedLetters := make([]uint64, bitsetSize)
			for i := 0; i < bitsetSize; i++ {
				newUsedLetters[i] = usedLetters[i] | pathInfo.lettersBitset[i]
			}
			dfs(index+1, append(selectedPaths, pathInfo.index), newUsedLetters)
		}
	}

	// Start DFS with empty selection
	usedLetters := make([]uint64, bitsetSize)
	dfs(0, []int{}, usedLetters)

	// Collect the best paths
	result := make([][]string, len(bestSelection))
	for i, idx := range bestSelection {
		result[i] = pathsInfo[idx].path
	}

	return result
}
