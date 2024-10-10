function dayOfTheYear(date) {

    const yearStart = new Date(date);
    yearStart.setDate(1)
    yearStart.setMonth(0)
    const dayOfYear = Math.round((date.getTime() - yearStart.getTime()) / (1000 * 60 * 60 * 24)) + 1;
    return dayOfYear;
}
