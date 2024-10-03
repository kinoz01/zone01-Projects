function triangle(str, n) {
    let r = ""
    for (let i = 1; i <= n; i++) {
        r += str.repeat(i) + "\n"
    }
    return r.slice(0, r.length-1)
}


console.log(triangle("s", 7))