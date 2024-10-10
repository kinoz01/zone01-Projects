function firstDayWeek(week, year) {
    if (week === 1) {
        return `01-01-${year}`
    }
    const msInDay = 1000 * 60 * 60 * 24; // Milliseconds in a day
    const startOfYear = new Date(`${year}-01-01`);
    let dayOfWeek = startOfYear.getDay();
    let offset = (dayOfWeek === 0 ? 6 : dayOfWeek - 1); 

    const firstMondayOfYear = new Date(startOfYear.getTime() - (offset * msInDay));

    const firstDayOfWeek = new Date(firstMondayOfYear.getTime() + (week - 1) * 7 * msInDay);

    if (firstDayOfWeek.getFullYear() < year) {
        return `01-01-${year}`;
    }
    
    const dd = String(firstDayOfWeek.getDate()).padStart(2, '0');
    const mm = String(firstDayOfWeek.getMonth() + 1).padStart(2, '0'); // Months are 0-based in JS
    const yyyy = String(firstDayOfWeek.getFullYear()).padStart(4, '0');

    return `${dd}-${mm}-${yyyy}`;
}

console.log(firstDayWeek(5, 1996))