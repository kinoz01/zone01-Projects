function debounce(func, wait) {
    let TimoutId
    return function (...args) {
        clearTimeout(TimoutId)
        TimoutId = setTimeout(() => func(...args), wait)
    }
}
// Its like we want to wrap a setTimeout with a reset that will
// rest the setTimeout each time the function is called
// TimoutId = setTimeout(() => func(...args), wait)
// VS
// TimoutId = setTimeout(func(...args), wait)
// in the second case func will run right away, and it's result  will be passed to setTimeout.
// setTimeout(func, wait) is expecting a function reference

function opDebounce(func, wait, leading) {
    let TimeoutId
    let leadingSwitch = false
    return function (...args) {
        const Immediately = !leadingSwitch && leading
        if (Immediately) {
            func(...args)
            leadingSwitch = true
        }
        clearTimeout(TimeoutId)
        TimeoutId = setTimeout(() => {
            if (!leading) func(...args)
            leadingSwitch = false
        }, wait)
    }
}

// Without leadingSwitch, debouncedLogger would log every single call immediately, ignoring the debounce mechanism entirely, making the leading behavior no different from simply executing the function immediately without any debounce.
// we need if (!leading) condition to prevent the execution of the function two times when leading is set to true. When leading is false we basically don't need this condition.

const logger = message => console.log(message);
const debouncedLogger = debounce(logger, 1000);

// Call debouncedLogger multiple times
// Each call reset the timer.
// Only the Third Call Is Printed after the set timeout.
debouncedLogger("First call");
debouncedLogger("Second call");
debouncedLogger("Third call");

// // debounce returns a new function that wraps the original func. 
// // Every time this returned function is called, it will debounce the execution of func.

// // If we don't use wrapper function we have to use timeoutId as a globale variable):
// let timeoutId;

// function debounce(func, wait, ...args) {
//     clearTimeout(timeoutId);
//     timeoutId = setTimeout(() => {
//         func(...args);
//     }, wait);
// }

// const logger1 = () => console.log("Logger 1");
// const logger2 = () => console.log("Logger 2");

// // Call debounce with two different functions
// debounce(logger1, 1000);  // Schedules logger1 to run in 1 second
// debounce(logger2, 1000)