
function isValid(date) {
    return (date !== 'Invalid Date' && !isNaN(date) && date !== '')
}

function isAfter(date1, date2) {
    return date1 > date2
}

function isBefore(date1, date2) {
    return date1 < date2
}

function isFuture(date) {
    return isValid(date) && date > Date.now()
}

function isPast(date) {
    return isValid(date) && date < Date.now()
}
