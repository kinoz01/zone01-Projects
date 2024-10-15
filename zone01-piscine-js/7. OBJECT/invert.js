function invert(obj) {
    let res = {}

    Object.keys(obj).forEach(key => {
        res[obj[key]] = key
    })
    return res
}