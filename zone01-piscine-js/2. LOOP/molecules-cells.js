function RNA(s) {
    let r = ""
    for (let i = 0; i < s.length; i++) {
        r += s[i] === 'G' ? 'C' :
            s[i] === 'C' ? 'G' :
            s[i] === 'T' ? 'A' :
            s[i] === 'A' ? 'U' :
            ''
    }
    return r;
}

function DNA(s) {
    let r = ""
    for (let i = 0; i < s.length; i++) {
        r += s[i] === 'C' ? 'G' :
            s[i] === 'G' ? 'C' :
            s[i] === 'A' ? 'T' :
            s[i] === 'U' ? 'A' :
            ''
    }
    return r;
}
