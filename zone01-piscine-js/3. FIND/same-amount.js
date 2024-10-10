function sameAmount(str, regex1, regex2) {
    let one = new RegExp(regex1.source, regex1.flags + "g");
    let two = new RegExp(regex2.source, regex2.flags + "g");

    const matches1 = str.match(one) || []  // If null, treat as empty array
    const matches2 = str.match(two) || []

    return matches1.length === matches2.length
}
