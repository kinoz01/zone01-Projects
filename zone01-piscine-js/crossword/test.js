const { crosswordSolver } = require('./crosswordSolver.js')



// Test 1:
const puzzle1 = `2001
0..0
1000
0..0`;
const words1 = ['casa', 'alan', 'ciao', 'anta'];

crosswordSolver(puzzle1, words1);
console.log("\n");
/*
Expected output:
casa
i..l
anta
o..n
*/

 
const puzzle2 = `...1...........
..1000001000...
...0....0......
.1......0...1..
.0....100000000
100000..0...0..
.0.....1001000.
.0.1....0.0....
.10000000.0....
.0.0......0....
.0.0.....100...
...0......0....
..........0....`
const words2 = [
    'sun',
    'sunglasses',
    'suncream',
    'swimming',
    'bikini',
    'beach',
    'icecream',
    'tan',
    'deckchair',
    'sand',
    'seaside',
    'sandals',
].reverse()

crosswordSolver(puzzle2, words2)
console.log("\n");

const puzzle3 = `..1.1..1...
10000..1000
..0.0..0...
..1000000..
..0.0..0...
1000..10000
..0.1..0...
....0..0...
..100000...
....0..0...
....0......`
const words3 = [
    'popcorn',
    'fruit',
    'flour',
    'chicken',
    'eggs',
    'vegetables',
    'pasta',
    'pork',
    'steak',
    'cheese',
]

crosswordSolver(puzzle3, words3)
console.log("\n");

const pu4 = '2001\n0..0\n2000\n0..0'
const words4 = ['casa', 'alan', 'ciao', 'anta']

crosswordSolver(pu4, words4)

const pu5 = '0001\n0..0\n3000\n0..0'
const words5 = ['casa', 'alan', 'ciao', 'anta']

crosswordSolver(pu5, words5)

const puz6 = '2001\n0..0\n1000\n0..0'
const words6 = ['casa', 'casa', 'ciao', 'anta']

crosswordSolver(puz6, words6)

const puz7 = '2001\n0..0\n1000\n0..0'
const words7 = ['aaab', 'aaac', 'aaad', 'aaae']

crosswordSolver(puz7, words7)
console.log("\n");

// Custom tests:
const puz8 = `...1...........
..1000001000...
...0....0......
.1......0...1..
.0....100000000
1000001.0...0..
.0....01001000.
.0.1..0.0.0....
.10000000.0...1
.0.0......0...0
.0.0.....100..0
...0......0...0
...1000...0...0`

const words8 = [
    "ammar",
    'what',
    'zer', 
    'sun',
    'sunglasses',
    'suncream',
    'swimming',
    'bikini',
    'beach',
    'icecream',
    'tan',
    'deckchair',
    'sand',
    'seaside',
    'sandals',
].reverse()

crosswordSolver(puz8, words8)
console.log("\n");


const puz9 = `...1...........
..1000001000...
...0....0......
.1......0...1..
.0....100000000
1000001.0...0..
.0....01001000.
.0.1..0.0.0....
.10000000.01...
.0.0......00...
.0.0.....100...
...0......0....
..........0....`

const words9 = [
    "zan",
    'zer', 
    'sun',
    'sunglasses',
    'suncream',
    'swimming',
    'bikini',
    'beach',
    'icecream',
    'tan',
    'deckchair',
    'sand',
    'seaside',
    'sandals',
].reverse()

crosswordSolver(puz9, words9)
console.log("\n");
