const map = (arr, func) => {
    let result = []
    for (let i = 0; i < arr.length; i++) {
        result.push(func(arr[i], i, arr))
    }
    return result
}

const flatMap = (arr, func) => {
    let result = []
    for (let i = 0; i < arr.length; i++) {
        result.push(func(arr[i], i, arr))
    }
    return [...result]
}