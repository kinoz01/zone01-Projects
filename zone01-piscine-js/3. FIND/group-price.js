function groupPrice(str) {
    const regex = /(USD|\$)(\d+)\.(\d+)/g
    
    let result = []
    const matches = [...str.matchAll(regex)];
    // console.log(matches);

    for (let match of matches) {
        let r = []
        r.push(match[0])
        r.push(match[2])
        r.push(match[3])
        result.push(r)
    }
    return result
}

console.log(groupPrice('The price USD1.2 of the cereals is 4.00 $9.98 and the second one its $10.20.'))