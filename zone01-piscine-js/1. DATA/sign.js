function sign(num) {
    return num==0 ? 0 : ( num>0 ? 1 : -1)  
}

function sameSign(n1, n2) {
    return (n1 > 0 && n2 > 0) || (n1 < 0 && n2 < 0) || (n1 === 0 && n2 === 0);
}

console.log(sameSign(-2, -1))