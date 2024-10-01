function isPositive(num) {
    return num>0 ? true : false
}

function abs(num) {
    return isPositive(num) ? num : -num
}

module.exports = { isPositive, abs };

