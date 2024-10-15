function fusion(...objs) {
    let res = {}
    objs.forEach(obj => {
        for (let key in obj) {
            if (!(key in res)) {
                res[key] = obj[key]
            } else {
                if (typeof obj[key] !== typeof res[key]) {
                    res[key] = obj[key]
                }
                else if (Array.isArray(obj[key]) && Array.isArray(res[key])) {
                    res[key] = res[key].concat(obj[key])
                }
                else if (typeof res[key] === 'string') {
                    res[key] += ' ' + obj[key]
                }
                else if (typeof res[key] === 'number') {
                    res[key] += obj[key]
                }
                else if (typeof res[key] === 'object' && !Array.isArray(res[key])) {
                    res[key] = fusion(res[key], obj[key])
                }
            }
        }
    })
    return res
}
