const currify = func => (...args) => args.length >= func.length ? func(...args) : (...more) => currify(func)(...args, ...more)


const add = (a, b, c) => a + b + c;
const curriedAdd = currify(add);
console.log(curriedAdd(1)(2)(3)); // Output: 6
console.log(curriedAdd(1, 2)(3)); // Output: 6
console.log(curriedAdd(1)(2, 3)); // Output: 6
