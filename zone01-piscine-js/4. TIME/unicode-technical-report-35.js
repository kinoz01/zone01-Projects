function format(date, fmt) {

    const year = date.getFullYear()
    const absYear = Math.abs(date.getFullYear())
    const month = date.getMonth() + 1
    const monthString = date.toLocaleString('default', { month: 'long' })
    const day = date.getDate()
    const dayString = date.toLocaleString('default', { weekday: 'long' });
    const minutes = date.getMinutes();
    let hours = date.getHours();
    let ampm = hours >= 12 ? 'PM' : 'AM';
    // Convert hours to 12-hour format
    let ampmHours = hours % 12;
    ampmHours = ampmHours ? ampmHours : 12; // If hours equals 0, make it 12
    const seconds = date.getSeconds();

    fmt = fmt.replace('a', ampm)

    fmt = fmt.replace('yyyy', String(absYear).padStart(4, '0'))
    fmt = fmt.replace('y', String(absYear))

    fmt = fmt.replace('MMMM', monthString)
    fmt = fmt.replace('MMM', monthString.slice(0, 3))
    fmt = fmt.replace('MM', String(month).padStart(2, '0'))
    fmt = fmt.replace(/^M$/g, String(month))
    fmt = fmt.replace(/(?<=\W)M(?=\W)/g, String(month))

    fmt = fmt.replace('dd', String(day).padStart(2, '0'))
    fmt = fmt.replace('d', String(day))

    fmt = fmt.replace('EEEE', String(dayString))
    fmt = fmt.replace('E', String(dayString).slice(0, 3))

    fmt = fmt.replace("mm", String(minutes).padStart(2, '0'))
    fmt = fmt.replace(/(?<=\W)m(?=\W)/g, String(minutes))

    fmt = fmt.replace('HH', String(hours).padStart(2, '0'))
    fmt = fmt.replace('H', String(hours))

    fmt = fmt.replace('hh', String(ampmHours).padStart(2, '0'))
    fmt = fmt.replace('h', String(ampmHours))

    fmt = fmt.replace('ss', String(seconds).padStart(2, '0'))
    fmt = fmt.replace('s', String(seconds))

    if (year <= 0) {
        fmt = fmt.replace('GGGG', 'Before Christ')
        fmt = fmt.replace('G', 'BC')
    } else {
        fmt = fmt.replace('GGGG', 'Anno Domini')
        fmt = fmt.replace('G', 'AD')
    }

    return fmt
}

// const landing = new Date('July 20, 1969, 20:17:40')
// const returning = new Date('July 21, 1969, 17:54:12')
// const eclipse = new Date(-585, 4, 28)
// const ending = new Date('2 September 1945, 9:02:14')

// // year
// console.log(format(eclipse, 'y'), '585')
// console.log(format(landing, 'y'), '1969')
// console.log(format(eclipse, 'yyyy'), '0585')
// console.log(format(landing, 'yyyy'), '1969')
// console.log(format(eclipse, 'yyyy G'), '0585 BC')
// console.log(format(landing, 'yyyy G'), '1969 AD')
// console.log(format(eclipse, 'yyyy GGGG'), '0585 Before Christ')
// console.log(format(landing, 'yyyy GGGG'), '1969 Anno Domini')

// // month
// console.log(format(eclipse, 'M'), '5')
// console.log(format(eclipse, 'MM'), '05')
// console.log(format(eclipse, 'MMM'), 'May')
// console.log(format(eclipse, 'MMMM'), 'May')
// console.log(format(landing, 'M'), '7')
// console.log(format(landing, 'MM'), '07')
// console.log(format(landing, 'MMM'), 'Jul')
// console.log(format(landing, 'MMMM'), 'July')
// console.log(format(ending, 'M'), '9')
// console.log(format(ending, 'MM'), '09')
// console.log(format(ending, 'MMM'), 'Sep')
// console.log(format(ending, 'MMMM'), 'September')

// // day
// console.log(format(landing, 'd'), '20')
// console.log(format(ending, 'd'), '2')
// console.log(format(landing, 'dd'), '20')
// console.log(format(ending, 'dd'), '02')
// console.log(format(landing, 'E'), 'Sun')
// console.log(format(returning, 'E'), 'Mon')
// console.log(format(landing, 'EEEE'), 'Sunday')
// console.log(format(returning, 'EEEE'), 'Monday')

// // time
// console.log(format(landing, 'H:m:s'), '20:17:40')
// console.log(format(landing, 'HH:mm:ss'), '20:17:40')
// console.log(format(landing, 'h:m:s a'), '8:17:40 PM')
// console.log(format(landing, 'hh:mm:ss a'), '08:17:40 PM')
// console.log(format(returning, 'H:m:s'), '17:54:12')
// console.log(format(returning, 'HH:mm:ss'), '17:54:12')
// console.log(format(returning, 'h:m:s a'), '5:54:12 PM')
// console.log(format(returning, 'hh:mm:ss a'), '05:54:12 PM')
// console.log(format(ending, 'H:m:s'), '9:2:14')
// console.log(format(ending, 'HH:mm:ss'), '09:02:14')
// console.log(format(ending, 'h:m:s a'), '9:2:14 AM')
// console.log(format(ending, 'hh:mm:ss a'), '09:02:14 AM')

// // mix
// console.log(format(ending, 'HH(mm)ss [dd] <MMM>'), '09(02)14 [02] <Sep>')
// console.log(format(ending, 'dd/MM/yyyy'), '02/09/1945')