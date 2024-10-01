const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the getAcceleration function from program.js
const { getAcceleration } = require('./physics.js');

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

// Test cases for getAcceleration function
t(({ eq }) => eq(getAcceleration({}), 'impossible'), "Should return 'impossible' when no valid inputs are provided");
t(({ eq }) => eq(getAcceleration({ d: 10, f: 2, Δv: 100 }), 'impossible'), "Should return 'impossible' when necessary inputs are missing");
t(({ eq }) => eq(getAcceleration({ f: 10, Δv: 100 }), 'impossible'), "Should return 'impossible' when mass is missing");
t(({ eq }) => eq(getAcceleration({ f: 10, m: 5 }), 2), "Should return 2 when force is 10 and mass is 5");
t(({ eq }) => eq(getAcceleration({ f: 10, m: 5, Δv: 100, Δt: 50 }), 2), "Should return 2 when both force/mass and Δv/Δt are provided");
t(({ eq }) => eq(getAcceleration({ Δv: 100, Δt: 50 }), 2), "Should return 2 when Δv is 100 and Δt is 50");
t(({ eq }) => eq(getAcceleration({ f: 10, Δv: 100, Δt: 50 }), 2), "Should return 2 when force, Δv, and Δt are provided");
t(({ eq }) => eq(getAcceleration({ f: 10, m: 5, Δt: 100 }), 2), "Should return 2 when force and mass are provided, ignoring Δt");
t(({ eq }) => eq(getAcceleration({ d: 10, t: 2, Δv: 100 }), 5), "Should return 5 when d is 10 and t is 2");
t(({ eq }) => eq(getAcceleration({ d: 100, t: 2, f: 100 }), 50), "Should return 50 when d is 100 and t is 2");

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
