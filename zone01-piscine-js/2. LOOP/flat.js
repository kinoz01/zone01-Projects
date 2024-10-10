function flat(arr, n = 1) {
    if (n === undefined) {
        return arr
    }
    let result = []

    if (n <= 0) {
        return arr
    }
    if (n == Infinity) {
        while (isNested(arr)) {
            result = flat1(arr)
            arr = result
        }
        return result
    }
    while (n > 0) {
        result = flat1(arr)
        arr = result
        n--
    }
    return result
}

const isNested = arr => arr.some(Array.isArray);
const flat1 = arr => arr.reduce((acc, val) => acc.concat(val), [])


// console.log(flat([1]));
// console.log(flat([1, [2, [3], [4, [5]]]], 3))
// console.log(flat([1, [2, [3], [4, [5]]]], Infinity));
// console.log(flat([1, [2, [3]]]));
// console.log(flat([1, [2, [3], [4, [5]]]], 2));