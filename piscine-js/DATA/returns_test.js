// test.js

// Import the functions from program.js
const { id, getLength } = require('./returns.js');

// Continue with the tests
const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

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

// Test cases for `id` function
t(() => eq(typeof id, 'function')); // id is declared and is a function
t(() => eq(id.length, 1)); // id takes 1 argument
t(() => eq(id(5), 5)); // id returns numbers back
t(() => eq(id('pouet'), 'pouet')); // id returns strings back
t(() => eq(id(id), id)); // id returns itself
t((_) => eq(id(_), _)); // id returns anything passed to it

// Test cases for `getLength` function
t(() => eq(getLength([2, 42]), 2)); // handle simple array
t(() => eq(getLength(['pouet', 4, true]), 3)); // handle mixed array
t(() => eq(getLength(Array(100)), 100)); // handle holey array
t(() => eq(getLength('salut'), 5)); // handle strings
t(() => eq(getLength([]), 0)); // handle empty arrays
t(() => eq(getLength(''), 0)); // handle empty strings

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
