function sign(num) {
    return num==0 ? 0 : ( num>0 ? 1 : -1)  
}

function sameSign(n1, n2) {
    return sign(n1)*sign(n2)>= 0
}

// Export the functions
module.exports = { sign, sameSign };
