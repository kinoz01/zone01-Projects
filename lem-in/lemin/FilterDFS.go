package lemin


func FilterPaths(paths [][]string, start, end string) [][]string {
    // Step 1: Precompute the set of unique strings in each path (excluding start and end)
    type PathInfo struct {
        index      int
        path       []string
        stringSet  map[string]bool
        length     int
    }

    pathsInfo := make([]PathInfo, len(paths))

    for i, path := range paths {
        stringSet := make(map[string]bool)
        for _, str := range path {
            if str != start && str != end {
                stringSet[str] = true
            }
        }
        pathsInfo[i] = PathInfo{
            index:     i,
            path:      path,
            stringSet: stringSet,
            length:    len(path),
        }
    }

    // Variables to store the best selection of paths
    var bestSelection []int
    maxNumPaths := 0
    minTotalLength := 0

    // Step 2: Implement DFS to explore path combinations
    var dfs func(pos int, selected []int, usedStrings map[string]bool)
    dfs = func(pos int, selected []int, usedStrings map[string]bool) {
        if pos == len(pathsInfo) {
            // Base case: all paths have been considered
            numSelected := len(selected)
            totalLength := 0
            for _, idx := range selected {
                totalLength += pathsInfo[idx].length
            }
            // Update best selection if better than current
            if numSelected > maxNumPaths || (numSelected == maxNumPaths && totalLength < minTotalLength) {
                maxNumPaths = numSelected
                minTotalLength = totalLength
                bestSelection = append([]int(nil), selected...)
            }
            return
        }

        // Prune if remaining paths can't improve the result
        remainingPaths := len(pathsInfo) - pos
        if len(selected)+remainingPaths < maxNumPaths {
            return
        }

        // Option 1: Exclude the current path
        dfs(pos+1, selected, usedStrings)

        // Option 2: Include the current path if no string conflict
        path := pathsInfo[pos]
        conflict := false
        for str := range path.stringSet {
            if usedStrings[str] {
                conflict = true
                break
            }
        }
        if !conflict {
            // Include current path and update used strings
            newUsedStrings := make(map[string]bool)
            for k, v := range usedStrings {
                newUsedStrings[k] = v
            }
            for str := range path.stringSet {
                newUsedStrings[str] = true
            }
            dfs(pos+1, append(selected, path.index), newUsedStrings)
        }
    }

    // Start DFS with empty selection and no used strings
    dfs(0, []int{}, make(map[string]bool))

    // Collect and return the best selection of paths
    result := make([][]string, len(bestSelection))
    for i, idx := range bestSelection {
        result[i] = pathsInfo[idx].path
    }

    return result
}
