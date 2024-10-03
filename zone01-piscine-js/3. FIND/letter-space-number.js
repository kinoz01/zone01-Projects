function letterSpaceNumber(str) {
    return str.match(/. \d((?=\W)|$)/g) || []
}

console.log(letterSpaceNumber('example 1, example 2'));