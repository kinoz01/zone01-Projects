// Return true if it found an input error.
function BasicChecks(puzzle, words) {
    // Check: 1. puzzle is a string, 2. words is an array,
    // 3. no repetition in words, 4. puzzle contains only /.012\n/
    if (
        typeof puzzle !== 'string' ||
        !Array.isArray(words) ||
        new Set(words).size !== words.length ||
        !/^[.012\n]+$/.test(puzzle)
    ) {
        return true;
    }

    // Check all words are strings
    for (let i = 0; i < words.length; i++) {
        if (typeof words[i] !== 'string') {
            return true;
        }
    }

    // Check puzzle is rectangular
    const rows = puzzle.split("\n");
    const firstRowLen = rows[0].length;
    for (let i = 1; i < rows.length; i++) {
        if (rows[i].length !== firstRowLen) {
            return true;
        }
    }

    // Check puzzle numbers match words numbers
    const c = [...puzzle].reduce((acc, char) => acc + (char == '1' ? 1 : char == '2' ? 2 : 0), 0);
    if (c !== words.length) {
        return true;
    }

    return false;
}

module.exports = { BasicChecks }