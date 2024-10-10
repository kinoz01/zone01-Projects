const vowels = /([aeiou])/gi

function vowelDots(str) {
    return str.replace(vowels, '$&.')
}

console.log(vowelDots('something'));