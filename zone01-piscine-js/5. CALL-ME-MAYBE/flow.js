const flow = funcs => (...args) => funcs.reduce((acc, func) => Array.isArray(acc) ? func(...acc) : func(acc), args)

function floww(funcs) {
    return function (...args) {
        let acc = funcs[0](...args)
        for (let i = 1; i < funcs.length; i++) {
            acc = funcs[i](acc)
        }
        return acc
    }
}

// const add = (x, y) => x + y;
// const square = x => x * x;
// const subtract = (x, y) => x - y;
// const double = x => x * 2;

// const flowFunc = floww([add, square, double]);

// console.log(flowFunc(2, 3))