export const generateLetters = () => {
    for (let i = 0; i < 120; i++) {
        let letter = document.createElement('div')
        letter.innerHTML = getLetter()
        letter.style.fontSize = `${i + 11}px`

        if (i < 40) letter.style.fontWeight = '300'
        else if (i < 80) letter.style.fontWeight = '400'
        else letter.style.fontWeight = '600'

        document.body.append(letter)
    }
}

export const getLetter = () => {
    const randomAscii = Math.floor(Math.random() * 26) + 65
    // Math.floor(Math.random() * 26) give a number between 0 (inclusive) and 26 (exclusive)
    return String.fromCharCode(randomAscii)
}