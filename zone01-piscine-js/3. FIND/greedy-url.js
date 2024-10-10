function getURL(str) {
    return str.match(/https?:\/\/\S+/g); 
}

function greedyQuery(str) {
    const regex = /https?:\/\/\S+\?\S+&\S+&\S+/g
    return str.match(regex)
}

function notSoGreedy(str) {
    let r =[]
    const regex = /https?:\/\/\S+\?(\S+)/g
    const matches = [...str.matchAll(regex)];
    for (let match of matches) {
        if (match[1].split('&').length==3 || match[1].split('&').length==2) {
            r.push(match[0])
        }
    }
    return r
}
