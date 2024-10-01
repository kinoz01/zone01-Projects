const escapeStr = "\`\\/\"\'"
console.log(escapeStr)
const arr = [4, '2']
const obj = {
    str: "hey",
    num: 42,
    bool: true,
    undef: undefined
};

const nested = Object.freeze({
    arr: Object.freeze([4, undefined, '2']),
    obj: Object.freeze({
        str: "hey",
        num: 42,
        bool: true
    })
});

module.exports = { escapeStr, arr, obj, nested }; // Export variables for testing
