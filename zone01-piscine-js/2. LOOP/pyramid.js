function pyramid(str, n) {
    let spacesN = 1
    let charN = 1
    let r = ""
    for (let i = 1; i < n; i++) {
        spacesN += 2*str.length
    }

    spacesN = (spacesN - 1) / 2

    for (let i = 0; i < n; i++) {
        r += " ".repeat(spacesN) + str.repeat(charN) + "\n"
        charN += 2
        spacesN -= str.length
    }

    return r.slice(0, r.length - 1)
}

console.log(pyramid("*", 7));
