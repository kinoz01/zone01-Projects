const neuron = (strArr) => {
    let result = {}

    strArr.forEach(str => {
        let [words, response] = str.split('/Response:/')
        words = words.split(/[:-]/)
        let type = words[0].toLowercase()
        if (!res[type]) res[type] = {}
        const elm = words[1].trim().split(' ').join('_').split(/\W/).join('').toLowerCase()
        if (!res[type][elm]) {
            res[type][elm] = {}
            res[type][elm][type.slice(0, -1)] = words[1].trim()
            res[type][elm]['responses'] = []
        }
        res[type][elm]['responses'].push(response.trim())
    });
}

console.log(neuron([
    'Questions: what is ounces? - Response: Ounce, unit of weight in the avoirdupois system',
    'Questions: what is ounces? - Response: equal to 1/16 pound (437 1/2 grains)',
    'Questions: what is Mud dauber - Response: Mud dauber is a name commonly applied to a number of wasps',
    'Orders: shutdown! - Response: Yes Sr!',
    'Orders: Quote something! - Response: Pursue what catches your heart, not what catches your eyes.'
  ]));