const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the `dogYears` function from dog-years.js
const { dogYears } = require('./dog-years.js');

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

// Test cases for `dogYears`
t(() => eq(dogYears('earth', 1000000000), 221.82));
t(() => eq(dogYears('mercury', 2134835688), 1966.16));
t(() => eq(dogYears('venus', 189839836), 68.45));
t(() => eq(dogYears('mars', 2129871239), 251.19));
t(() => eq(dogYears('jupiter', 901876382), 16.86));
t(() => eq(dogYears('saturn', 2000000000), 15.07));
t(() => eq(dogYears('uranus', 1210123456), 3.19));
t(() => eq(dogYears('neptune', 1821023456), 2.45));

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
