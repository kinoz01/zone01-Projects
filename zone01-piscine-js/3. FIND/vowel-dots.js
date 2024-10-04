function vowelDots(str) {
    const vowels = /([aeiou])/gi
    return str.replcae(vowels, '$&.')
}

console.log(vowelDots('something'));