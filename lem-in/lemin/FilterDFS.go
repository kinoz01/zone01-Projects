package lemin

func GenerateNonCrossingPathGroups(paths [][]string, start, end string) [][][]string {
    // Step 1: Precompute the set of unique strings in each path (excluding start and end)
    type PathInfo struct {
        path      []string
        stringSet map[string]bool
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
            path:      path,
            stringSet: stringSet,
        }
    }

    // Step 2: Use DFS to find all combinations of non-crossing paths
    var result [][][]string

    var dfs func(index int, currentGroup [][]string, usedStrings map[string]bool)
    dfs = func(index int, currentGroup [][]string, usedStrings map[string]bool) {
        if index == len(pathsInfo) {
            if len(currentGroup) > 0 {
                // Add a copy of the current group to the result
                groupCopy := make([][]string, len(currentGroup))
                copy(groupCopy, currentGroup)
                result = append(result, groupCopy)
            }
            return
        }

        // Option 1: Exclude the current path and move to the next
        dfs(index+1, currentGroup, usedStrings)

        // Option 2: Include the current path if no string conflict
        pathInfo := pathsInfo[index]
        conflict := false
        for str := range pathInfo.stringSet {
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
            for str := range pathInfo.stringSet {
                newUsedStrings[str] = true
            }
            dfs(index+1, append(currentGroup, pathInfo.path), newUsedStrings)
        }
    }

    // Start DFS with empty group and no used strings
    dfs(0, [][]string{}, make(map[string]bool))

    return result
}
