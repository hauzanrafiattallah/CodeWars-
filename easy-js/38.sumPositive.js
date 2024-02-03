// https://www.codewars.com/kata/5715eaedb436cf5606000381/train/javascript

function positiveSum(arr) {
  var sum = 0;
  for (var i = 0; i < arr.length; i++) {
    if (arr[i] > 0) {
      sum += arr[i];
    }
  }
  return sum;
}

// Contoh penggunaan:
var numbers = [1, -4, 7, 10000];
var result = positiveSum(numbers);
console.log(result); // Output: 20
