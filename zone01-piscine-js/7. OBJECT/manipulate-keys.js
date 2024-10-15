function filterKeys(obj, func) {
    let res = {}
    Object.keys(obj).forEach(key => {
        if (func(key)) res[key] = obj[key]
    })
    return res
}

function mapKeys(obj, func) {
    let res = {}
    Object.keys(obj).forEach(key => {
        res[func(key)] = obj[key]
    })
    return res
}

// we use here ... to pass initial value and a function
function reduceKeys(obj, ...funcAndInit) {
    return Object.keys(obj).reduce(...funcAndInit)
}
