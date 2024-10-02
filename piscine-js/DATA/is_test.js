const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the `is` object from the relevant module
const { is } = require('./is.js');

// Test setup
const tests = [];
const t = (f) => tests.push(f);

// eq function to compare results
const eq = (actual, expected) => {
    if (JSON.stringify(actual) === JSON.stringify(expected)) {
        console.log(`${colors.green}Test passed:${colors.reset} Test ${JSON.stringify(actual)} === ${JSON.stringify(expected)}${colors.reset}`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} ${JSON.stringify(actual)} !== ${JSON.stringify(expected)}${colors.reset}`);
    }
};

// Setup function returns the context array for testing
const setup = () => [
    0,
    NaN,
    true,
    '',
    '💩',
    undefined,
    t,
    [],
    {},
    [1, Array(1), [], 2],
    { length: 10 },
    Object.create(null),
    null,
    console.log,
    void 0,
];

// Match function filters the values for the function being tested
const match = ({ eq, ctx }, fun, values) => eq(ctx.filter(fun), values);

// Testing each `is` function
t(() => match({ eq, ctx: setup() }, is.num, [0, NaN]));
t(() => match({ eq, ctx: setup() }, is.nan, [NaN]));
t(() => match({ eq, ctx: setup() }, is.str, ['', '💩']));
t(() => match({ eq, ctx: setup() }, is.bool, [true]));
t(() => match({ eq, ctx: setup() }, is.undef, [undefined, undefined]));
t(() => match({ eq, ctx: setup() }, is.arr, [[], [1, Array(1), [], 2]]));
t(() => match({ eq, ctx: setup() }, is.obj, [{}, { length: 10 }, Object.create(null)]));
t(() => match({ eq, ctx: setup() }, is.fun, [t, console.log]));
t(() => match({ eq, ctx: setup() }, is.falsy, [0, NaN, '', undefined, null, void 0]));

// is.def tests
t(() => !setup().filter(is.def).includes(undefined));
t(() => setup().filter(is.def).length === setup().length - 2);

// is.truthy tests
t(() => match({ eq, ctx: setup() }, is.truthy, [
    true,
    '💩',
    t,
    [],
    {},
    [1, Array(1), [], 2],
    { length: 10 },
    Object.create(null),
    console.log,
]));

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
