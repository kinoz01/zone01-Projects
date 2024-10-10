function sums(n) {
    function partition(num, min) {
        if (num === 0) {
            return [[]]
        }

        let result = []
        for (let i = min; i <= num; i++) {
            let subPartitions = partition(num - i, i)
            for (let sub of subPartitions) {
                result.push([i, ...sub])
            }
        }
        return result
    }
    return partition(n, 1)
}

// Example usage:
console.log(sums(4));