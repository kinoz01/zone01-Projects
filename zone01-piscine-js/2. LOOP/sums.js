function sums(n) {
    function partition(number, min) {

        if (number === 0) {
            return [[]]; // Base case: If number is 0, return an empty partition
        }

        let result = [];
        for (let i = min; i <= number; i++) {
            const subPartitions = partition(number - i, i);
            subPartitions.forEach(sub => {
                result.push([i, ...sub]); 
            });
        }
        return result;
    }

    return partition(n, 1).slice(0,-1);
}

// Example usage:
console.log(sums(4));