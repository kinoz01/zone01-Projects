function indexOf(arr, val, index) {
    if (!index) {
        index = 0
    }
    for (let i = index; i < arr.length; i++) {
        if (arr[i] === val) {
            return i
        }
    }
    return -1
}

function lastIndexOf(arr, val, index) {
    if (!index) {
        index = arr.length - 1
    }
    for (let i = index; i >= 0; i--) {
        if (arr[i] === val) {
            return i
        }
    }
    return -1
}

function includes(arr, val) {
    if (indexOf(arr, val) !== -1) {
        return true
    }
    return false
}
