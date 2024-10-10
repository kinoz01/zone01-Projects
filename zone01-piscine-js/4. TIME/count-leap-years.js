function countLeapYears(date) {
    let year = date.getFullYear()
    let c = 0
    for (let i = 1; i < year; i++) {
        if (i%100 === 0 && i%400!==0) {
            continue
        }
        if (i % 4 === 0) {
            c++
        }
    }
    return c
}

console.log(countLeapYears(new Date('0001-12-01')) === 0)
console.log(countLeapYears(new Date('1664-08-09')) === 403)
console.log(countLeapYears(new Date('2020-01-01')) === 489)
console.log(countLeapYears(new Date('2048-12-08')) === 496)