const sourceObject = {
    num: 42,
    bool: true,
    str: 'some text',
    log: console.log,
}

function get(key) {
    return sourceObject[key]
}

function set(key, val) {
    sourceObject[key] = val
    return val
}

module.exports = { get, set };
