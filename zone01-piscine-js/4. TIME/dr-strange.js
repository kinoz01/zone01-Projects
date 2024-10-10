function addWeek(date) {
    const daysOfWeek = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday', 'secondMonday', 'secondTuesday', 'secondWednesday', 'secondThursday', 'secondFriday', 'secondSaturday', 'secondSunday',]
    const start = new Date('0001-01-01')
    const D = new Date(date)
    const days = Math.round((D.getTime() - start.getTime()) / (1000 * 60 * 60 * 24))
    return daysOfWeek[days % 14]
}

function timeTravel({date, hour, minute, second}) {
    date.setHours(hour)
    date.setMinutes(minute)
    date.setSeconds(second)
    return date
}
