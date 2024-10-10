function matchCron(cron, date) {
    const dateSlice = [date.getMinutes(), date.getHours(), date.getDate(), date.getMonth() + 1, date.getDay()]

    const part = cron.split(" ")

    for (let i = 0; i < part.length; i++) {       
        if (part[i] === '*') {
            continue
        }
        if (dateSlice[i] !== Number(part[i])) {
            return false
        }
    }
    return true
}


console.log(matchCron('* * * * 1', new Date('2020-06-01 00:00:00')))
console.log(matchCron('* * * 2 *', new Date('2021-02-01 00:00:00')))
console.log(matchCron('* * 9 * *', new Date('2020-06-09 00:00:00')))
console.log(matchCron('* 3 * * *', new Date('2020-05-31 03:00:00')))
console.log(matchCron('1 * * * *', new Date('2020-05-30 19:01:00')))
console.log(matchCron('3 3 * 3 3', new Date('2021-03-03 03:03:00')))
console.log(matchCron('* * * * *', new Date()))

console.log(!matchCron('* * * * 1', new Date('2020-06-02 00:00:00')))
console.log(!matchCron('* * * 2 *', new Date('2021-03-01 00:00:00')))
console.log(!matchCron('* * 8 * *', new Date('2020-06-09 00:00:00')))
console.log(!matchCron('* 2 * * *', new Date('2020-05-31 03:00:00')))
console.log(!matchCron('1 * * * *', new Date('2020-05-30 19:00:00')))
console.log(!matchCron('3 3 * 3 3', new Date('2021-03-02 03:03:00')))
