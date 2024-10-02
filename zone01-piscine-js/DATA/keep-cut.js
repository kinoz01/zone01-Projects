const cutFirst = str => str.slice(2)
const cutLast = str => str.slice(0, -2);
const cutFirstLast = str => str.slice(2, -2);
const keepFirst = str => str.slice(0, 2);
const keepLast = str => str.slice(-2);
const keepFirstLast = str => str.length <= 4 ? str : str.slice(0, 2) + str.slice(-2);

console.log(keepFirstLast('af'));