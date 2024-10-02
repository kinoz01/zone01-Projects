function reverse(input) {
    let res= []
    for (let i =input.length -1; i>=0; i--) {
        res.push(input[i])
    }
    return typeof input === "string" ? res.join('') : res
}