
const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the `concatStr` function from `program.js`
const { concatStr } = require('./concat-str.js');

// Test setup
const tests = [];
const t = (f, message) => {
    try {
        f();
        console.log(`${colors.green}Test passed:${colors.reset} ${message}`);
    } catch (err) {
        console.error(`${colors.red}Test failed:${colors.reset} ${message}`);
    }
};

// Test cases for `concatStr`
t(() => typeof concatStr === 'function', 'Should be a function');
t(() => concatStr.length === 2, 'Should take 2 arguments');
t(() => concatStr('a', 'b') === 'ab', 'Concatenates two strings correctly');
t(() => concatStr('yolo', 'swag') === 'yoloswag', 'Concatenates "yolo" and "swag" correctly');

// Handle non-string inputs correctly
t(() => concatStr(1, 2) === '12', 'Handles non-strings correctly (numbers)');
t(() => concatStr(concatStr, concatStr) === String(concatStr).repeat(2), 'Handles non-strings correctly (functions)');

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
