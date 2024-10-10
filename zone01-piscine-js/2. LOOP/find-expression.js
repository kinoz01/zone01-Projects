const findExpression = (num) => {

    let result = "1 "
    let resultint = 1
    let twoC = 0
    let fourC = 0

    if (num === 1) {
        return "1"
    }
    if (num % 2 != 0 || num < 1) {
        return undefined
    }


    while ((num - resultint) % 4 != 0) {
        resultint = 2 * resultint
        twoC++
    }

    while (resultint < num) {
        resultint = resultint + 4
        fourC++
    }
    result += repeat(mul2 + " ", twoC) + repeat(add4 + " ", fourC)
    
    return result.slice(0, -1)
}

const repeat = (str, n) => n <= 0 ? "" : str + repeat(str, n - 1)

console.log(findExpression(4));