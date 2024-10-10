function citiesOnly(cities) {
    return cities.map(city => city.city)
}

function upperCasingStates(states) {
    return states.map(state => state.split(" ").map(part => part.charAt(0).toUpperCase() + part.slice(1)).join(" "))
}

function fahrenheitToCelsius(fTemps) {
    return fTemps.map(temp => Math.floor((Number(temp.slice(0, -2)) - 32) * 5 / 9) + "°C")
}

function trimTemp(cities) {
    return cities.map(city => ({ ...city, temperature: city.temperature.replace(/\s+/g, '') }))
}

function tempForecasts(cities) {
    return cities.map(city => `${Math.floor((Number(city.temperature.replace(/\s+/g, '').slice(0, -2)) - 32) * 5 / 9)}°Celsius in ${city.city}, ${city.state.split(" ").map(cityPart => cityPart.charAt(0).toUpperCase() + cityPart.slice(1)).join(" ")}`)
}

console.log(tempForecasts([
    {
      city: 'Pasadena',
      temperature: ' 101 °F',
      state: 'california dfdf',
      region: 'West',
    },
  ]));