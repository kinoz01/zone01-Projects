// ANSI color codes
const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import variables from program.js
const { escapeStr, arr, obj, nested } = require('./declarations');

// Test setup
const tests = [];
const t = (f) => tests.push(f);
const isConst = (name) => {
    try {
        eval(`${name} = 'm'`);
        return false;
    } catch (err) {
        return true;
    }
};

// Helper function to check for failed assignments
const cantEdit = (fn) => {
    try {
        fn();
    } catch (err) {
        return true;
    }
};

// eq function to compare values
const eq = (actual, expected) => {
    if (JSON.stringify(actual) === JSON.stringify(expected)) {
        console.log(`${colors.green}Test passed:${colors.reset} ${JSON.stringify(actual)} === ${JSON.stringify(expected)}`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} ${JSON.stringify(actual)} !== ${JSON.stringify(expected)}`);
    }
};

// Tests
t(() => typeof escapeStr === 'string');  // escapeStr is declared and of type string
t(() => escapeStr.includes("'"));        // should include the character '
t(() => escapeStr.includes('"'));        // should include the character "
t(() => escapeStr.includes('`'));        // should include the character `
t(() => escapeStr.includes('/'));        // should include the character /
t(() => escapeStr.includes('\\'));       // should include the character \

t(() => Array.isArray(arr));  // arr is declared and is an array
t(() => eq(arr[0], 4));       // arr first element is 4
t(() => eq(arr[1], '2'));     // arr second element is "2"
t(() => eq(arr.length, 2));   // arr length is 2

t(() => obj.constructor === Object);    // obj is declared and of type object
t(() => typeof obj.str === 'string');   // obj.str is of type string
t(() => typeof obj.num === 'number');   // obj.num is of type number
t(() => typeof obj.bool === 'boolean'); // obj.bool is of type boolean
t(() => 'undef' in obj && typeof obj.undef === 'undefined');  // obj.undef is of type undefined

t(() => nested.constructor === Object);  // nested is declared and is of type object
t(() => eq(nested.arr[0], 4));           // nested.arr first element is 4
t(() => eq(nested.arr[1], undefined));   // nested.arr second element is undefined
t(() => eq(nested.arr[2], '2'));         // nested.arr third element is "2"
t(() => eq(nested.arr.length, 3));       // nested.arr length is 3

t(() => cantEdit(() => (nested.obj = 5)));        // nested.obj is frozen and can not be reassigned
t(() => cantEdit(() => (nested.arr.push('test')))); // nested.arr is not frozen and can be modified
t(() => eq(nested.arr.length, 3));              // nested.arr length remains 3

// Check if all variables are constant
t(() => ['escapeStr', 'arr', 'obj', 'nested']
    .every(isConst));

// Run all tests
Object.freeze(tests);
tests.forEach(test => test({ eq }));
