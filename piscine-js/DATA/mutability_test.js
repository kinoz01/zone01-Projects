const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

const person = {
    name: 'Rick',
    age: 77,
    country: 'US',
};

module.exports = { person };

// Import the functions and variables from your main program
const { samePerson, clone1, clone2 } = require('./mutability.js');

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

// Test cases with eq
t(() => eq(typeof samePerson, 'object'));
t(() => eq(typeof person, 'object'));
t(() => eq(typeof clone1, 'object'));
t(() => eq(typeof clone2, 'object'));
t(() => eq(clone1, clone2)); // Equal values
t(() => eq(person.name, 'Rick'));
t(() => eq(person.age, 78)); // Incremented age
t(() => eq(person.country, 'FR')); // Modified country
t(() => eq(clone1.country, 'US')); // Clone1 remains unchanged
t(() => eq(clone2.age, 77)); // Clone2 remains unchanged

// Handling reference checks directly
t(() => {
    if (clone1 !== clone2) {
        console.log(`${colors.green}Test passed:${colors.reset} clone1 !== clone2`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} clone1 === clone2`);
    }
});

t(() => {
    if (person === samePerson) {
        console.log(`${colors.green}Test passed:${colors.reset} person === samePerson`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} person !== samePerson`);
    }
});

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
