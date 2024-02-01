// https://www.codewars.com/kata/55d24f55d7dd296eb9000030/train/javascript

function summation(num) {
  let result = 0;
  for (let i = 1; i <= num; i++) {
    result += i;
  }
  return result;
}

// Example usage:
console.log(summation(2)); // Output: 3 (1 + 2)
console.log(summation(8)); // Output: 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)
