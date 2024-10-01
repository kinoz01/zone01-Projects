const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the `biggie` and `smalls` variables from program.js
const { biggie, smalls } = require('./biggie-smalls.js');

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

// Test cases for `biggie` and `smalls`
t(() => typeof biggie !== 'undefined', 'biggie should be defined');
t(() => biggie > 1.7976931348623157e308, 'biggie should be larger than the maximum double-precision floating-point number');

t(() => typeof smalls !== 'undefined', 'smalls should be defined');
t(() => smalls < -1.7976931348623157e308, 'smalls should be smaller than the minimum double-precision floating-point number');

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());

/*
  “Damn right I like the life I live,
   because I went from negative to positive.”

      ― The Notorious B.I.G
*/
