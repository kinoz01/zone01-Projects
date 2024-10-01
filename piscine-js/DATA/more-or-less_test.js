const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the functions
const { more, less, add, sub } = require('./more-or-less.js');

// Test setup
const tests = [];
const t = (f) => tests.push(f);

// eq function to compare results
const eq = (actual, expected) => {
    if (actual === expected) {
        console.log(`${colors.green}Test passed:${colors.reset} Test ${JSON.stringify(actual)} === ${JSON.stringify(expected)}${colors.reset}`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} ${JSON.stringify(actual)} !== ${JSON.stringify(expected)}${colors.reset}`);
    }
};

// Test cases
t(() => eq(typeof more, 'function'));
t(() => eq(more.length, 1));
t(() => eq(more(5), 6));
t(() => eq(more(more(more(5))), 8)); // more more more !!

t(() => eq(typeof less, 'function'));
t(() => eq(less.length, 1));
t(() => eq(less(5), 4));
t(() => eq(less(1), more(-1)));

t(() => eq(typeof add, 'function'));
t(() => eq(add.length, 2));
t(() => eq(add(3, 10), 13));
t(() => eq(add(-1, -1), -2));

t(() => eq(typeof sub, 'function'));
t(() => eq(sub.length, 2));
t(() => eq(sub(-1, -1), 0));
t(() => eq(sub(3, 10), -7));

// Combination of functions
t(() => eq(add(more(10), sub(less(5), 10)), 5));

const rand = Math.random();
t(() => eq(less(rand), rand - 1));
t(() => eq(more(rand), rand + 1));

// Freeze the tests array to prevent further modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
