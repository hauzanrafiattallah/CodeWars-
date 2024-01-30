// https://www.codewars.com/kata/5168bb5dfe9a00b126000018

function solution(str) {
  // Step 1: Convert the string to an array of characters
  var charArray = str.split("");

  // Step 2: Reverse the array
  var reversedArray = charArray.reverse();

  // Step 3: Join the array back into a string
  var reversedString = reversedArray.join("");

  // Return the reversed string
  return reversedString;
}

// Example usage:
console.log(solution("hauzan")); // Output: "nazuah"
