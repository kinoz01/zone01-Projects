function replica(...objs) {
    let res = {}
    objs.forEach(obj => {
        Object.keys(obj).forEach(key => {
            if (typeof obj[key] === 'object' && typeof res[key] === 'object'
                && !Array.isArray(obj[key]) && !Array.isArray(res[key])) {
                res[key] = { ...res[key], ...obj[key] }
            } else res[key] = obj[key]
        })
    })
    return res
}

const obj1 = {
    a: 1,
    b: { x: 10, y: 20 } // spread merge
};

const obj2 = {
    b: { y: 25, z: 30 }, // spread merge
    c: 5
};

console.log(replica(obj1, obj2));
