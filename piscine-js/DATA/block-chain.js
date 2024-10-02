function blockchain(data, prev) {
    if (prev == null) {
        prev = {
            index: 0,
            hash: '0'
        }   
    }
    const index = prev.index + 1
    const DataBlock = JSON.stringify(data)
    
}