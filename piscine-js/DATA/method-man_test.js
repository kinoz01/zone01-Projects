const colors = {
    green: '\x1b[32m',
    red: '\x1b[31m',
    reset: '\x1b[0m',
};

// Import the functions from method-man.js
const { words, sentence, yell, whisper, capitalize } = require('./method-man.js');

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

// Function to generate a random string
const randomString = () => Math.random().toString(36).substring(7);

// Generate a random string to be used in tests
const r = randomString();

// Test cases for `words` function
t(() => eq(words('a b c'), ['a', 'b', 'c']), "Words should split by space");
t(() => eq(words('Hello  world'), ['Hello', '', 'world']), "Words should handle double spaces");
t(() => eq(words(`${r} ${r} ${r}`), [r, r, r]), "Words should split repeated random words");

// Test cases for `sentence` function
t(() => eq(sentence(['a', 'b', 'c']), 'a b c'), "Sentence should join words with spaces");
t(() => eq(sentence(['Hello', '', 'world']), 'Hello  world'), "Sentence should handle empty strings between words");
t(() => eq(sentence([r, r, r]), `${r} ${r} ${r}`), "Sentence should join repeated random words");

// Test cases for `yell` function
t(() => eq(yell('howdy stranger ?'), 'HOWDY STRANGER ?'), "Yell should uppercase all letters");
t(() => eq(yell('Déjà vu'), 'DÉJÀ VU'), "Yell should uppercase special characters");

// Test cases for `whisper` function
t(() => eq(whisper('DÉJÀ VU'), '*déjà vu*'), "Whisper should lowercase all letters and wrap with *");
t(() => eq(whisper('HOWDY STRANGER ?'), '*howdy stranger ?*'), "Whisper should lowercase and wrap with *");

// Test cases for `capitalize` function
t(() => eq(capitalize('str'), 'Str'), "Capitalize should capitalize first letter and lowercase others");
t(() => eq(capitalize('qsdqsdqsd'), 'Qsdqsdqsd'), "Capitalize should handle lowercase input");
t(() => eq(capitalize('STR'), 'Str'), "Capitalize should handle uppercase input");
t(() => eq(capitalize('zapZAP'), 'Zapzap'), "Capitalize should handle mixed case");
t(() => eq(capitalize('zap ZAP'), 'Zap zap'), "Capitalize should handle words with spaces");

// Freeze the tests array to prevent modification
Object.freeze(tests);

// Run all tests
tests.forEach(test => test());
