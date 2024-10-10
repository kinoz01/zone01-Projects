const adder = (numArr, initVal = 0) => numArr.reduce((c, n) => c += n, initVal)

const sumOrMul = (numArr, initVal = 0) => numArr.reduce((c, n) => n % 2 === 0 ? c * n : c + n, initVal)

const funcExec = (funcArr, num = 0) => funcArr.reduce((c, func) => func(c), num)
