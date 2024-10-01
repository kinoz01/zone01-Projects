const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the `is` object from program.js
const { is } = require('./is.js');

// Test setup
const tests = [];
const t = (f, description) => {
    try {
        f();
        console.log(`${colors.green}Test passed:${colors.reset} ${description}`);
    } catch (err) {
        console.error(`${colors.red}Test failed:${colors.reset} ${description}`);
    }
};

// eq function to compare results
const eq = (actual, expected) => {
    if (JSON.stringify(actual) === JSON.stringify(expected)) {
        return true;
    } else {
        throw new Error(`${actual} !== ${expected}`);
    }
};

// Sample context array to use in tests
const ctx = [
    0,
    NaN,
    '',
    '💩',
    true,
    [],
    [1, Array(1), [], 2],
    { length: 10 },
    Object.create(null),
    null,
    console.log,
    void 0,
];

// Helper function to match values for `is` functions
const match = (fun, values) => {
    const truthyValues = ctx.filter(fun);
    const expected = values;
    const others = ctx.filter(val => !values.includes(val));

    return eq(truthyValues, expected) && others.every(val => !fun(val));
};

// Test cases for `is` functions
t(() => match(is.num, [0, NaN]), "is.num should return true for numbers, including NaN, and false for others");
t(() => match(is.nan, [NaN]), "is.nan should return true for NaN and false for others");
t(() => match(is.str, ['', '💩']), "is.str should return true for strings and false for others");
t(() => match(is.bool, [true]), "is.bool should return true for booleans and false for others");
t(() => match(is.undef, [void 0]), "is.undef should return true for undefined values and false for others");
t(() => match(is.arr, [[], [1, Array(1), [], 2]]), "is.arr should return true for arrays and false for others");
t(() => match(is.obj, [{}, { length: 10 }, Object.create(null)]), "is.obj should return true for objects and false for others");
t(() => match(is.fun, [t, console.log]), "is.fun should return true for functions and false for others");
t(() => match(is.falsy, [0, NaN, '', undefined, null, void 0]), "is.falsy should return true for falsy values and false for others");

// is.def
t(() => ctx.filter(is.def).length === ctx.length - 2, "is.def should return true for defined values except undefined");

// is.truthy
t(() => match(is.truthy, [
    true,
    '💩',
    t,
    [],
    {},
    [1, Array(1), [], 2],
    { length: 10 },
    Object.create(null),
    console.log,
]), "is.truthy should return true for truthy values and false for others");

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
