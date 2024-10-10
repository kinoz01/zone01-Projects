// Find words places and return an array with:
// col, row, maxVLen, maxHLen and possible directions (H, V, H or V)

function WordsPlacesFinder(GRID, shortest, X, Y) {
    let WordsPlaces = [];
    for (let i = 0; i < Y; i++) {
        for (let j = 0; j < X; j++) {

            if (GRID[i][j] == '1') {
                // we can put word only horizontally
                if (j === X || (GRID[i][j + 1] === '.')) {
                    let k = i;
                    while (k < Y && GRID[k][j] !== '.') {
                        k++;
                    }
                    if (k - i >= shortest) {
                        WordsPlaces.push({ row: i, col: j, maxVLen: k - i, maxHLen: 0, direction: 'V' });
                    }
                } else {
                    // We can put word both horizontally OR vertically
                    if (i + 1 != Y && GRID[i + 1][j] !== '.') {
                        let x = j;
                        let y = i;
                        while (y < Y && GRID[y][j] !== '.') {
                            y++;
                        }
                        while (x < X && GRID[i][x] !== '.') {
                            x++;
                        }
                        if (x - j >= shortest && y - i >= shortest) {
                            WordsPlaces.push({ row: i, col: j, maxVLen: y - i, maxHLen: x - j, direction: 'H or V' });
                        } else if (x - j >= shortest) {
                            WordsPlaces.push({ row: i, col: j, maxVLen: 0, maxHLen: x - j, direction: 'H' });
                        } else if (y - i >= shortest) {
                            WordsPlaces.push({ row: i, col: j, maxVLen: y - i, maxHLen: 0, direction: 'V' });
                        }

                    } else { // We can put word only horizontally
                        let m = j;
                        while (m < X && GRID[i][m] !== '.') {
                            m++;
                        }
                        if (m - j >= shortest) {
                            WordsPlaces.push({ row: i, col: j, maxVLen: 0, maxHLen: m - j, direction: 'H' });
                        }
                    }
                }
            }
            if (GRID[i][j] === '2') {
                let m = j;
                let n = i;
                while (n < Y && GRID[n][j] !== '.') {
                    n++;
                }
                while (m < X && GRID[i][m] !== '.') {
                    m++;
                }
                if (m - j >= shortest) {
                    WordsPlaces.push({ row: i, col: j, maxVLen: 0, maxHLen: m - j, direction: 'H' });
                }
                if (n - i >= shortest) {
                    WordsPlaces.push({ row: i, col: j, maxVLen: n - i, maxHLen: 0, direction: 'V' });
                }
            }
        }
    }
    return WordsPlaces;
}

module.exports = { WordsPlacesFinder };