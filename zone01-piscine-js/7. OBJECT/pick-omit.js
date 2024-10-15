function pick(obj, str) {
    if (typeof str === 'string') str = [str]
    let res = {}
    str.forEach(key => {
        if (obj[key] !== undefined) res[key] = obj[key]
    })
    return res
}

function omit(obj, str) {
    let res = {}
    if (typeof str === 'string') str = [str]
    Object.keys(obj).forEach(key => {
        if (!str.includes(key)) res[key] = obj[key]
    })
    
    return res
}

/* Object.keys(obj): This method returns an array of an object's 
own enumerable properties (i.e., properties that belong directly 
to the object and are not inherited from the object's prototype). 
*/

// using delete and deep copy for omit
/*
function omit(obj, str) {
    let res = clone(obj)
    if (typeof str === 'string') str = [str]
    str.forEach(key => {
        delete res[key]
    })
    return res
}

function clone(obj) {
    if (obj === null || typeof obj !== 'object') return obj
    if (Array.isArray(obj)) return obj.map(item => clone(item))
    const res = {}
    for (let key in obj) {
        if (obj.hasOwnProperty(key)) {
            res[key] = clone(obj[key])
        }
    }
    return res
}
*/