function split(str, sep) {
    let index = 0
    let res = []

    if (sep === '') {
        for (let i = 0; i < str.length; i++) {
            res.push(str[i]);
        }
        return res;
    }

    for (let i = 0; i <= str.length - sep.length;) {
        if (str.slice(i, i + sep.length) === sep) {
            res.push(str.slice(index, i))
            index = i + sep.length
            i = index
        } else {
            i++
        }
    }
    res.push(str.slice(index))
    return res
}

function join(arr, char) {
    let str = ""
    for (let i = 0; i < arr.length; i++) {
        str += arr[i]
        if (i !== arr.length - 1) {
            str += char
        }
    }
    return str
}

const str = "hey    ho w are you  k  "
console.log(split(str, " "));

console.log(split('a b c', ' '))
console.log(split('ggg - ddd - b', ' - '))
console.log(split('ee,ff,g,', ','));
console.log(split('Riad', ' '));
console.log(split('rrrr', 'rr'));
console.log(split('rrirr', 'rr'));
console.log(split('Riad', ''));
console.log(split('', 'Riad'));