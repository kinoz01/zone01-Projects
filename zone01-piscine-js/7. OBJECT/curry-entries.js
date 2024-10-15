function defaultCurry(obj1) {
    return function (obj2) {
        return Object.fromEntries(Object.entries(obj1).concat(Object.entries(obj2)))
        // if there are duplicate keys, the last occurrence of 
        // the key will overwrite the previous one
    }
}

const defCurry = obj1 => obj2 => Object.fromEntries(Object.entries(obj1).concat(Object.entries(obj2)))

function mapCurry(func) {
    return function (obj) {
        return Object.fromEntries(Object.entries(obj).map(func))
    }
}

const reduceCurry = func => (obj, init) => Object.entries(obj).reduce(func, init)

const filterCurry = func => obj => Object.fromEntries(Object.entries(obj).filter(func))

function reduceScore(obj, init = 0) {
    return reduceCurry((acc, [, val]) => {
        return val.isForceUser ? acc + val.pilotingScore + val.shootingScore : acc
    })(obj, init)
}

function filterForce(obj) {
    return filterCurry(([, val]) => {
        return val.isForceUser && val.shootingScore >= 80
    })(obj)
}

function mapAverage(obj) {
    return mapCurry(([key, val]) => {
        val.averageScore = (val.pilotingScore + val.shootingScore) / 2
        return [key, val]
    })(obj)
}


// const personnel = {
//     lukeSkywalker: { id: 5, pilotingScore: 98, shootingScore: 56, isForceUser: true },
//     sabineWren: { id: 82, pilotingScore: 73, shootingScore: 99, isForceUser: false },
//     zebOrellios: { id: 22, pilotingScore: 20, shootingScore: 59, isForceUser: false },
//     ezraBridger: { id: 15, pilotingScore: 43, shootingScore: 67, isForceUser: true },
//     calebDume: { id: 11, pilotingScore: 71, shootingScore: 85, isForceUser: true },
// }

// console.log(mapCurry(([k, v]) => [`${k}_force`, v])(personnel))

// console.log(reduceCurry((acc, [k, v]) => (acc += v))({ a: 1, b: 2, c: 3 }, 0));
// console.log(defaultCurry({
//     http: 403,
//     connection: 'close',
//     contentType: 'multipart/form-data',
// })({
//     http: 200,
//     connection: 'open',
//     requestMethod: 'GET'
// }));