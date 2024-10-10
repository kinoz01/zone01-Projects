const every = (arr, func) => arr.filter(func).length === arr.length
const some = (arr, func) => arr.filter(func).length > 0
const none = (arr, func) => arr.filter(func).length === 0