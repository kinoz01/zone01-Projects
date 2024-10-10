const filterShortStateName = arr => arr.filter(str => str.length < 7)

const filterStartVowel = arr => arr.filter(str => /^[aeiou]/i.test(str))

const filter5Vowels = arr => arr.filter(str => [...str].reduce((acc, val) => /^[aeiou]$/i.test(String(val)) ? acc + 1 : acc, 0) >= 5)

const filter1DistinctVowel = arr => arr.filter(str => [...str].reduce((acc, val) => /^[aeiou]$/i.test(val) ? acc += val.toLowerCase() : acc, '').split('').every((char, _, arr) => char === arr[0]))

const multiFilter = arrObjects => arrObjects.filter(obj => obj.capital.length >= 8 && !/^[aeiou]/i.test(obj.name.charAt(0)) && /[aeiou]/i.test(obj.tag) && obj.region !== "South") 