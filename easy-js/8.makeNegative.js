// https://www.codewars.com/kata/55685cd7ad70877c23000102

function makeNegative(num) {
  if (num <= 0) {
    return num;
  } else {
    return num * -1;
  }
}

// Example usage:
console.log(makeNegative(-9)); // Output: -5
console.log(makeNegative(0)); // Output: 0
