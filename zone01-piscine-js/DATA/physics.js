function getAcceleration(newtonObj) {
    return newtonObj.f !== undefined && newtonObj.m !== undefined
    ? newtonObj.f / newtonObj.m
    : newtonObj.Δv !== undefined && newtonObj.Δt !== undefined
    ? newtonObj.Δv / newtonObj.Δt
    : newtonObj.d !== undefined && newtonObj.t !== undefined
    ? (2 * newtonObj.d) / (newtonObj.t ** 2)
    : "impossible"
} 
