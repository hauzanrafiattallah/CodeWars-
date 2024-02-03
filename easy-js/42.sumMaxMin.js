// https://www.codewars.com/kata/576b93db1129fcf2200001e6/train/javascript

function sumArray(array) {
  if (!array || array.length <= 1) {
    return 0;
  }

  let min = array[0];
  let max = array[0];
  let sum = 0;

  for (let i = 0; i < array.length; i++) {
    sum += array[i];

    if (array[i] < min) {
      min = array[i];
    }

    if (array[i] > max) {
      max = array[i];
    }
  }

  return sum - min - max;
}

const arr1 = [6, 2, 1, 8, 10];
const result1 = sumArray(arr1);
console.log(result1); // Output: 16
