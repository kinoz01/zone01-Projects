function chunk(arr , c) {
    let temp = []
    let res = []

    for (let i = 0; i< arr.length; i++) {
        if (i%c == 0 && i!=0) {
            res.push(temp)
            temp = []
        }
        temp.push(arr[i])
    }
    res.push(temp)
    return res
}
