function vowelDots(str) {
    const vowels = /([aeiou])/gi
    return str.replace(vowels, '$&.')
}

console.log(vowelDots('something'));