const arrToSet = arr => new Set(arr);
const arrToStr = arr => arr.join('');
const setToArr = set => [...set];
const setToStr = set => [...set].join('');
const strToArr = str => [...str];
const strToSet = str => new Set(str);
const mapToObj = map => {
    let obj = {}
    for (let [key, value] of map) {
        obj[key] = value;
    }
    return obj
};
const objToArr = obj => Object.values(obj);
const objToMap = obj => new Map(Object.entries(obj));
const arrToObj = arr => arr.reduce((acc, val, idx) => ({ ...acc, [idx]: val }), {});
const strToObj = str => [...str].reduce((acc, char, idx) => ({ ...acc, [idx]: char }), {});

const superTypeOf = val => {
    if (val === null) return 'null';
    if (Array.isArray(val)) return 'Array';
    if (val instanceof Set) return 'Set';
    if (val instanceof Map) return 'Map';
    if (val == undefined) {
        return 'undefined'
    }
    const type = typeof val;
    return type.charAt(0).toUpperCase() + type.slice(1);
};

console.log(superTypeOf(undefined))