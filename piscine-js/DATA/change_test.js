const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the `get` and `set` functions from `program.js`
const { get, set } = require('./change');

// Test setup
const tests = [];
const t = (f) => tests.push(f);

// eq function to compare results using your preferred format
const eq = (actual, expected) => {
    if (actual === expected) {
        console.log(`${colors.green}Test passed:${colors.reset} Test ${JSON.stringify(actual)} === ${JSON.stringify(expected)}${colors.reset}`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} ${JSON.stringify(actual)} !== ${JSON.stringify(expected)}${colors.reset}`);
    }
};

// Test cases for `get` function
t(() => eq(typeof get, 'function')); // get is a function
t(() => eq(get('num'), 42)); // get retrieves 'num' correctly
t(() => eq(get('bool'), true)); // get retrieves 'bool' correctly
t(() => eq(get('str'), 'some text')); // get retrieves 'str' correctly
t(() => eq(get('log'), console.log)); // get retrieves 'log' correctly
t(() => eq(get('noexist'), undefined)); // get returns undefined for non-existent keys

// Test cases for `set` function
t(() => eq(typeof set, 'function')); // set is a function
t(() => eq(set('num', 55), 55)); // set updates 'num' and returns new value
t(() => eq(set('noexist', 'nice'), 'nice')); // set adds new key 'noexist' and returns value
t(() => eq(get('num'), 55)); // get retrieves updated 'num' value
t(() => eq(get('noexist'), 'nice')); // get retrieves newly added key 'noexist'
t(() => eq(set('log', undefined), undefined)); // set 'log' to undefined
t(() => eq(get('log'), undefined)); // get retrieves 'log' after being set to undefined

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
