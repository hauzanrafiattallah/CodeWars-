// https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/train/javascript

function invert(array) {
  let result = [];
  for (let i = 0; i < array.length; i++) {
    result.push(-array[i]);
  }
  return result;
}
console.log(invert([-1, 2, 3, 0]));
