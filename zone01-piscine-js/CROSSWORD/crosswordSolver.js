const { BasicChecks } = require('./checkPuzzle.js')
const { WordsPlacesFinder } = require('./findWords.js')
const { fillGrid } = require('./fillPuzzle.js') 

function crosswordSolver(puzzle, words) {
    // Check for input errors
    if (BasicChecks(puzzle, words)) {
        console.log('Error');
        return;
    }

    let GRID = puzzle.split("\n").map(row => row.split(''))
    const Y = GRID.length;
    const X = GRID[0].length;

    // Sort words from longest to shortest
    words.sort((a, b) => a.length - b.length);
    const shortest = words[0].length;

    let wordPlaces = WordsPlacesFinder(GRID, shortest, X, Y);

    let solutions = [];
    let usedWords = new Array(words.length).fill(false);

    fillGrid(GRID, words, wordPlaces, 0, usedWords, solutions)

    if (solutions.length === 1) {
        // Only one unique solution found
        // console.log(solutions[0])
        const result = solutions[0].map(row => row.join('')).join('\n');
        console.log(result);
    } else {
        // No solution or multiple solutions
        console.log('Error');
    }
}

module.exports = { crosswordSolver }
