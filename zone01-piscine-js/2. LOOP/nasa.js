function nasa(N) {
    let s = ""
    for (let i =1; i <=N; i++) {
        if (i%3 == 0 && i%5==0) {
            s+=  "NASA "
        } else if (i%3 == 0 ) {
            s+= "NA "
        } else if (i%5 == 0) {
            s += "SA "
        } else{
            s += String(i) + " "
        }
        
    }
    return s.slice(0, s.length-1)
}

