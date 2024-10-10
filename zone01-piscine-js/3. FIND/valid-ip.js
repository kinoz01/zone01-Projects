function findIP(str) {

    let result = []
    const regex = /(\d+)\.(\d+)\.(\d+)\.(\d+)(:\d+)?/g

    const matches = [...str.matchAll(regex)];
    for (let match of matches) {
        let leadingZeros = false

        for (let j = 1; j <= 4; j++) {
            if (match[j].slice(0, 1) === '0' && match[j] != 0) {
                leadingZeros = true
            }
            if (match[j] == 0 && match[j].length>1) {
                leadingZeros = true
            }
        }
        if (!leadingZeros && match[1] < 256 && match[2] < 256 && match[3] < 256 && match[4] < 256) {
            if (match[5]) {
                if (match[5].slice(1) <= 65535) {
                    result.push(match[0])
                }
            } else {
                result.push(match[0])
            }
        }
    }
    return result
}
