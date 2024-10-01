const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Disable the built-in Math.min and Math.max
Math.min = Math.max = undefined;

// Import the functions from program.js
const { max, min } = require('./min-max');

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

// Test cases for `max` function
t(() => max(0, -2) === 0, "max should return the larger of 0 and -2");
t(() => max(-1, 10) === 10, "max should return the larger of -1 and 10");
t(() => max(-13.2, -222) === -13.2, "max should return the larger of -13.2 and -222");
t(() => max(132, 133) === 133, "max should return the larger of 132 and 133");

// Test cases for `min` function
t(() => min(0, -2) === -2, "min should return the smaller of 0 and -2");
t(() => min(-1, 10) === -1, "min should return the smaller of -1 and 10");
t(() => min(-13.2, -222) === -222, "min should return the smaller of -13.2 and -222");
t(() => min(132, 133) === 132, "min should return the smaller of 132 and 133");

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
