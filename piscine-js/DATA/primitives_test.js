const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

const { str, num, bool, undef } = require('./primitives.js');

// Helper function to check if a variable is declared as a constant
const isConst = (name) => {
  try {
    eval(`${name} = 'modified'`);  // Try to modify the variable
    return false; // If modification succeeds, it's not a constant
  } catch (err) {
    return true;  // If an error occurs (due to const), it is a constant
  }
};

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
t(() => eq(typeof str, 'string'));   // str is declared and of type string
t(() => eq(typeof num, 'number'));   // num is declared and of type number
t(() => eq(typeof bool, 'boolean')); // bool is declared and of type boolean
t(() => eq(undef, undefined));       // undef is declared and of type undefined

// Check if all variables are const
t(() => {
    const areConst = ['str', 'num', 'bool', 'undef'].every(isConst);
    if (areConst) {
        console.log(`${colors.green}Test passed:${colors.reset} All variables are const`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} Not all variables are const`);
    }
});

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
