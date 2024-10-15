function deepCopy(obj) {
    if (obj === null || typeof obj !== 'object') return obj

    if (Array.isArray(obj)) return obj.map(cell => deepCopy(cell))

    let res = {}
    Object.keys(obj).forEach(key => {
        res[key] = deepCopy(obj[key])
    })
    return res
}