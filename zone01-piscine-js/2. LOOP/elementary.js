const multiply = (a, b) => {
    return b === 0
        ? 0
        : (b > 0
            ? a + multiply(a, b - 1)
            : -multiply(a, -b))
}

const divide = (a, b) => {
    if (b === 0) {
        return NaN
    }

    const sign = (a < 0) !== (b < 0) ? -1 : 1
    let up = Math.abs(a)
    const down = Math.abs(b)
    let r = 0

    while (up >= down) {
        up -= down
        r++
    }
    return sign == -1 ? -r : r
}

function modulo(a, b) {
    if (b === 0) return NaN
    const pa = Math.abs(a)
    const pb = Math.abs(b)

    return a < 0
        ? -modulo(-a, b)
        : pa < pb
            ? a
            : modulo(pa - pb, b)
}
