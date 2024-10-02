const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Disable the built-in Math.sign
Math.sign = undefined;

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

// Test cases for `sign` function
t(() => typeof sign === 'function', "sign should be a function");
t(() => sign.length === 1, "sign should take 1 argument");
t(() => sign !== Math.sign, "sign should not be the same as Math.sign");
t(() => sign(-2) === -1, "sign should return -1 for negative numbers");
t(() => sign(10) === 1, "sign should return 1 for positive numbers");
t(() => sign(0) === 0, "sign should return 0 for zero");
t(() => sign(132) === 1, "sign should return 1 for large positive numbers");

// Test cases for `sameSign` function
t(() => typeof sameSign === 'function', "sameSign should be a function");
t(() => sameSign.length === 2, "sameSign should take 2 arguments");
t(() => sameSign(-2, -1) === true, "sameSign should return true for two negative numbers");
t(() => sameSign(0, 0) === true, "sameSign should return true for two zero values");
t(() => sameSign(12, 3232) === true, "sameSign should return true for two positive numbers");
t(() => sameSign(1, -1) === false, "sameSign should return false for a positive and negative number");
t(() => sameSign(-231, 1) === false, "sameSign should return false for a negative and positive number");
t(() => sameSign(-231, 0) === false, "sameSign should return false for a negative number and zero");
t(() => sameSign(0, 231) === false, "sameSign should return false for zero and a positive number");
t(() => sameSign(231, -233) === false, "sameSign should return false for a positive and negative number");

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
