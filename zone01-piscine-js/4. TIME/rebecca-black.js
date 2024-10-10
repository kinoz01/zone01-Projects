function isFriday(date) {
    return date.getDay() == 5
}

function isWeekend(date) {
    return (date.getDay() == 0 || date.getDay() == 6)
}

function isLeapYear(date) {
    return date.getFullYear() % 4 === 0
}

function isLastDayOfMonth(date) {
    const nextDay = new Date(date)
    nextDay.setDate(date.getDate() + 1)
    return nextDay.getMonth() !== date.getMonth()
}
