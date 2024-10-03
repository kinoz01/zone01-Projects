function ionOut(str) {
    return str.match(/(\S+t)(?=ion)/gi) || []
}

// golang do not support lookahead assertions:
function inOUtlong(str) {
    let r =[]
    const regex = /(\S+t)(ion)/g
    const matches = [...str.matchAll(regex)]

    for (let match of matches) {
        r.push(match[1])
    }
    return r
}

console.log(ionOut('1st position is the vision of the 2nd position'));
