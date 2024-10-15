function filterValues(obj, func) {
    let res = {}
    Object.keys(obj).forEach(key => {
        if (func(obj[key])) res[key] = obj[key]
    })
    return res
}

function mapValues(obj, func) {
    let res = {}
    Object.keys(obj).forEach(key => {
        res[key] = func(obj[key])
    })
    return res
}

function reduceValues(obj, func, res = 0) {
    return Object.keys(obj).reduce((acc, key) => func(acc, obj[key]), res)
}
