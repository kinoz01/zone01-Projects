function throttle(func, rate) {
    let lastExc = 0 // CLOSER variable
    return function() {
        let now = Date.now()
        
        if (now-lastExc > rate) {
            func(...args)
            lastExc = now
        }

    }
}

// Example callback function
const log = (arr, el) => arr.push(el);

// Example usage
const throttledLog = throttle(log, 100);

let arr = [];
let interval = setInterval(() => throttledLog(arr, 1), 50);
setTimeout(() => clearInterval(interval), 250);
