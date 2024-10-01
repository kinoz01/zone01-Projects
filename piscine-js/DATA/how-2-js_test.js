// ANSI color codes
const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

const { readFileSync: read } = require('fs'); // Use require instead of import

const tests = [
  ({ eq }) =>
    eq(
      read('./index.html', 'utf8').trim(),
      '<script type="module" src="how-2-js.js"></script>',
    ),
  ({ eq }) =>
    eq(
      read('./how-2-js.js', 'utf8').trim(),
      `console.log('Hello World')`,
    ),
];

const eq = (actual, expected) => {
    if (JSON.stringify(actual) === JSON.stringify(expected)) {
        console.log(`${colors.green}Test passed:${colors.reset} ${JSON.stringify(actual)} === ${JSON.stringify(expected)}`);
    } else {
        console.error(`${colors.red}Test failed:${colors.reset} ${JSON.stringify(actual)} !== ${JSON.stringify(expected)}`);
    }
};

// Run tests
tests.forEach((test) => test({ eq }));
