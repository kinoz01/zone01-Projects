const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the `circular` object from program.js
const { circular } = require('./circular.js');

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

// Test cases for `circular` object
t(() => circular.constructor === Object, 'circular should be an object');
t(() => circular.circular === circular, 'circular.circular should reference circular itself');
t(() => circular.circular.circular === circular, 'circular.circular.circular should reference circular itself');
t(() => circular.circular.circular.circular === circular, 'circular.circular.circular.circular should reference circular itself');
t(() => circular.circular.circular.circular.circular === circular, 'circular.circular.circular.circular.circular should reference circular itself');

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
