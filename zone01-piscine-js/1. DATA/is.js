is.num = val => typeof val == 'number';
is.nan = val => Number.isNaN(val);
is.str = val => typeof val == 'string';
is.bool = val => typeof val === 'boolean';
is.undef = val => typeof val === 'undefined';
is.def = val => typeof val !== 'undefined';
is.arr = val => Array.isArray(val);
is.obj = val => val !== null && typeof val === 'object' && !Array.isArray(val);
is.fun = val => typeof val === 'function';
is.truthy = val => !!val;
is.falsy = val => !val;



// Traditional function
/* function double(n) {
    return n * 2;
} */

// Arrow function with parentheses, curly braces, and return
/* const double = (n) => {
    return n * 2;
}; */

// const double = (n) => n * 2;

// Shortened arrow function (no curly braces, implicit return)
// const double = n => n * 2;