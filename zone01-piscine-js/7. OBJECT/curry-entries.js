function defaultCurry(obj1, obj2) {
    let res = {}
    Object.keys(obj1).forEach(key => {
        res[key] = obj1[key]
    })
    Object.keys(obj2).forEach(key => {
        res[key] = obj2[key]
    })
    return res
}

console.log(defaultCurry({
    http: 403,
    connection: 'close',
    contentType: 'multipart/form-data',
  })({
    http: 200,
    connection: 'open',
    requestMethod: 'GET'
  }));