function sameAmount(str, regex1, regex2) {
    let one = new RegExp(regex1.source, regex1.flags + "g");
    let two = new RegExp(regex2.source, regex2.flags + "g");

    const matches1 = str.match(one) || []  // If null, treat as empty array
    const matches2 = str.match(two) || []

    return matches1.length === matches2.length
}

const data = `qqqqqqq q qqqqqqqfsqqqqq q qq  qw w wq wqw  wqw
ijnjjnfapsdbjnkfsdiqw klfsdjn fs fsdnjnkfsdjnk sfdjn fsp fd`

console.log(sameAmount('hello how are you', /h/, /e/))
console.log(sameAmount('hello how are you', /he/, /ho/))
console.log(sameAmount('hello how are you', /l/, /e/))
console.log(sameAmount(data, /i/, /p/))
console.log(sameAmount(data, /h/, /w/) )
console.log(sameAmount(data, /qqqq /, /qqqqqqq/))
console.log(sameAmount(data, /q /, /qqqqqqq/))
console.log(sameAmount(data, /fs[^q]/, /q /))
console.log(sameAmount(data, /^[qs]/, /^[gq]/))
console.log(sameAmount(data, /j/, /w/))
console.log(sameAmount(data, /j/, / /))
console.log(sameAmount(data, /fs./, /jn./))