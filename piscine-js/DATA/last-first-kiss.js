function first(strArr) {
    return strArr[0]
}

function last(strArr) {
    return strArr[strArr.length-1]
}

function kiss(strArr) {
    return [last(strArr), first(strArr)]
}

// Exporting the functions using CommonJS syntax
module.exports = { first, last, kiss };