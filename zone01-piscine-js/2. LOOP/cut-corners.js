function round(num) {
    if (num >= 0) {
        let x = modulo(num * 10, 10)
        let y = x / 10
        num -= y
        if (x>=5) {
            return num+1
        }
        return num
    } else {
        let x = modulo(-num * 10, 10)
        let y = x / 10
        num += y
        if (x>=5) {
            return num-1
        }
        return num
    }
}

function ceil(num) {
    if (num == 0) {
        return 0
    }
    if (num > 0) {
        let x = modulo(num * 10, 10)
        let y = x / 10
        num -= y
        return num+1
    } else {
        let x = modulo(-num * 10, 10)
        let y = x / 10
        num += y
        return num
    }
}

function floor(num) {
    if (num >= 0) {
        let x = modulo(num * 10, 10)
        let y = x / 10
        num -= y
        return num
    } else {
        let x = modulo(-num * 10, 10)
        let y = x / 10
        num += y
        return num-1
    }
}

function trunc(num) {
    if (num >= 0) {
        let x = modulo(num * 10, 10)
        let y = x / 10
        num -= y
        return num
    } else {
        let x = modulo(-num * 10, 10)
        let y = x / 10
        num += y
        return num
    }
}

function modulo(a, b) {
    if (b === 0) return NaN

    const pa = Math.abs(a)
    const pb = Math.abs(b)

    let result = pa
    while (result >= pb) {
        result -= pb
    }

    return a < 0 ? -result : result
}

// 
// Math.PI, -Math.PI, Math.E, -Math.E, 0
// nums.map(floor), [3, -3, 3, -3, 0])
// [4, -3, 3, -2, 0])
console.log(ceil(Math.PI));
console.log(ceil(-Math.PI));
console.log(ceil(Math.E));
console.log(ceil(-Math.E));
console.log(ceil(0));
