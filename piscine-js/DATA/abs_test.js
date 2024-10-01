const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Disable the built-in Math.abs
Math.abs = undefined;

// Import the functions from program.js
const { isPositive, abs } = require('./abs.js');

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

// Test cases for `isPositive` function
t(() => isPositive(3) === true, "isPositive should return true for positive numbers");
t(() => isPositive(1998790) === true, "isPositive should return true for large positive numbers");
t(() => isPositive(-1) === false, "isPositive should return false for negative numbers");
t(() => isPositive(-0.7) === false, "isPositive should return false for negative decimals");
t(() => isPositive(-787823) === false, "isPositive should return false for large negative numbers");
t(() => isPositive(0) === false, "isPositive should return false for zero");

// Test cases for `abs` function
t(() => eq(abs(0), 0), "abs should return 0 for input 0");
t(() => eq(abs(-1), 1), "abs should return positive 1 for input -1");
t(() => eq(abs(-13.2), 13.2), "abs should return 13.2 for input -13.2");
t(() => eq(abs(132), 132), "abs should return 132 for input 132");

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
