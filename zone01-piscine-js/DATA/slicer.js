function slice(arr, start, end = arr.length) {
    let r = []
    start < 0 ? start += arr.length : start
    end < 0 ? end += arr.length : end

    if (start >= end) {
        return typeof arr == "string" ? "" : r
    }
    for (let i = start; i < end; i++) {
        typeof arr == "string" ? r += arr[i] : r.push(arr[i])
    }

    return r
}

// don't need to handle:
console.log(slice("hello", 2));