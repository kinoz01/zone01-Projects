// fill the grid using backtracking to find a solution

function fillGrid(GRID, words, wordPlaces, index, usedWords, solutions) {

    if (solutions.length > 1) {
        return // More than one solution found
    }
    if (index === wordPlaces.length) {
        // All words placed
        // Deep copy the grid to save the solution
        console.log("hey");
        solutions.push(GRID.map(row => row.slice()))
        return
    }
    let place = wordPlaces[index]
    let possibleDirections = []
    if (place.direction.includes('H')) {
        possibleDirections.push('H')
    }
    if (place.direction.includes('V')) {
        possibleDirections.push('V')
    }

    for (let direction of possibleDirections) {
        let maxLen = direction === 'H' ? place.maxHLen : place.maxVLen

        for (let i = 0; i < words.length; i++) {
            if (usedWords[i]) continue

            let word = words[i]
            if (word.length > maxLen) continue
            if (canPlaceWord(GRID, word, place, direction)) {
                let original = placeWord(GRID, word, place, direction)
                usedWords[i] = true

                fillGrid(GRID, words, wordPlaces, index + 1, usedWords, solutions);

                // Backtrack
                removeWord(GRID, original, place, direction);
                usedWords[i] = false; // Unmark word
            }
        }
    }
}

function canPlaceWord(GRID, word, place, direction) {
    // cell -----> GRID[y][x]
    // col ------> x
    // row ------> y
    let y = place.row;
    let x = place.col;
    if (direction === 'H') {
        if (x + word.length > GRID[0].length) {
            return false // Exceeds grid horizontally
        }
        for (let i = 0; i < word.length; i++) {
            let cell = GRID[y][x + i]
            if (cell === '.' || (cell !== '0' && cell !== '1' && cell != '2' && cell != word[i])) {
                return false
            }
        }
    } else if (direction === 'V') {
        if (y + word.length > GRID.length) {
            return false; // Exceeds grid vertically
        }
        for (let i = 0; i < word.length; i++) {
            let cell = GRID[y + i][x]
            if (cell === '.' || (cell !== '0' && cell !== '1' && cell !== '2' && cell !== word[i])) {
                return false
            }
        }
    }
    return true
}

function placeWord(GRID, word, place, direction) {
    let { row, col } = place;
    let original = [];
    if (direction === 'H') {
        for (let i = 0; i < word.length; i++) {
            original.push(GRID[row][col + i]);
            GRID[row][col + i] = word[i];
        }
    } else if (direction === 'V') {
        for (let i = 0; i < word.length; i++) {
            original.push(GRID[row + i][col]);
            GRID[row + i][col] = word[i];
        }
    }
    return original;
}

function removeWord(GRID, original, place, direction) {
    console.log(original);
    console.log(GRID);
    let { row, col } = place;
    if (direction === 'H') {
        for (let i = 0; i < original.length; i++) {
            GRID[row][col + i] = original[i];
        }
    } else if (direction === 'V') {
        for (let i = 0; i < original.length; i++) {
            GRID[row + i][col] = original[i];
        }
    }  
}

module.exports = { fillGrid }