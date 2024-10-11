const currify = func => (...args) => args.length >= func.length ? func(...args) : (...more) => currify(func)(...args, ...more)

function currifyy(func) {
    return function(...args) {
        if (args.length >= func.length) {
           return func(...args) // Call and return the function when enough arguments are gathered
        } else {
            return function(...nextArgs) {
                return currify(func)(...args, ...nextArgs) // Return the recursive call to gather more arguments
            }
        }
    }
}

// const add = (a, b, c) => a + b + c;
// const curriedAdd = currify(add);
// const curried = currify(curriedAdd(1)) // return of the first call of currify

// console.log(curried(2)(3)); // Output: 6
// console.log(curriedAdd(1, 2)(3)); // Output: 6
// console.log(curriedAdd(1)(2, 3)); // Output: 6
