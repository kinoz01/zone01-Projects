import { colors } from './fifty-shades-of-cold.data.js'

export const generateClasses = () => {
    let cold = document.createElement('style')

    for (let color of colors) {
        cold.innerHTML += `.${color} {background: ${color};}`
    }
    document.head.append(cold)
}

export const generateColdShades = () => {
    for (let color of colors) {
        if ((/(aqua|blue|turquoise|green|cyan|navy|purple)/).test(color)) {
            let coldColor = document.createElement('div')
            coldColor.classList.add(color)
            coldColor.textContent = color
            document.body.append(coldColor)
        }
    }
    
}

export const choseShade = (shade) => {
    let coldColrs = Array.from(document.getElementsByTagName('div'))
    for (let color of coldColrs) {
        color.className = ''
        color.classList.add(shade)
    }
}