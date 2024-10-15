function pronoun(str) {
    const words = str.split(/[\s,]+/)
    let res = {}
    for (let i = 0; i < words.length; i++) {
        words[i] = words[i].toLowerCase()
        if (isPronoun(words[i])) {
            if (!res[words[i]]) {
                res[words[i]] = { word: [], count: 0 }
            }
            if (i !== words.length - 1 && !isPronoun(words[i + 1])) {
                res[words[i]].word.push(words[i+1])
            }
            res[words[i]].count++
        }
    }
    return res
}

function isPronoun(str) {
    const pronouns = ['i', 'you', 'he', 'she', 'it', 'we', 'they']
    for (const key in pronouns) {
        if (pronouns[key].toLowerCase() === str.toLowerCase()) return true
    }
    return false
}