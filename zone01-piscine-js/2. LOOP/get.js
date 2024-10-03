function get(src, path) {
    let arr = path.split(".")

    for (let key of arr ) {
        src = src[key]
        if (src=== undefined) {
            return undefined
        }
    }
    return src
}

console.log(get({ a: [{ b: t }] }, 'a.0.b.toString'));