function dogYears(planet, dogAge) {
    const StoD = 60*60*24
    const DtoY = {
        earth: 1.0,
        mercury: 0.2408467,
        venus: 0.61519726,
        mars: 1.8808158, 
        jupiter: 11.862615,
        saturn: 29.447498, 
        uranus: 84.016846,
        neptune: 164.79132
    }

    const DaytoYear = DtoY[planet];
    if (!DaytoYear) {
        console.log('Invalid planet name');
    }

    return Math.round((dogAge/(StoD*365.25*DaytoYear))*7*100)/100
}

module.exports = { dogYears };