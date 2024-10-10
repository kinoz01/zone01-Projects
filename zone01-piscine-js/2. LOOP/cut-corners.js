const ceil = (n) => {
    if (get(n) == 0) {
      return n;
    }
    if (n < 0) {
      return n + get(n);
    }
    return n - get(n) + 1;
  };
  
  const floor = (n) => {
    if (get(n) == 0) {
      return n;
    }
    if (n < 0) {
      return n + get(n) - 1;
    }
    return n - get(n);
  };
  
  const round = (n) => {
    if (n < 0) {
      let a = get(n);
      if (a * -1 >= -0.5) {
        return trunc(n);
      }
      return trunc(n) - 1;
    }
    if (get(n) >= 0.5) {
      return trunc(n) + 1;
    }
    return trunc(n);
  };
  const trunc = (a) => {
    if (a > -1 && a < 1) {
      return 0;
    }
    if (a < 0) {
      return a + get(a);
    }
    return a - get(a);
  };
  const get = (a) => {
    if (a >= 0xfffffffff) {
      a = a - 0xfffffffff;
    }
    if (a > -1 && a < 1) {
      return a;
    }
    if (a < 0) {
      a = a * -1;
    }
    while (a / 10 < 0) {
      a /= 10;
    }
  
    while (a > 0) {
      if (a - 1 < 0) {
        break;
      }
      a -= 1;
      console.log(a);
    }
    return a;
  };

// Math.PI, -Math.PI, Math.E, -Math.E, 0
// nums.map(floor), [3, -3, 3, -3, 0])
// [4, -3, 3, -2, 0])
console.log(ceil(Math.PI));
console.log(ceil(-Math.PI));
console.log(ceil(Math.E));
console.log(ceil(-Math.E));
console.log(ceil(0));
