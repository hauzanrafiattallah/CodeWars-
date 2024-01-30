// https://www.codewars.com/kata/56f695399400f5d9ef000af5/train/javascript

function correctTail(body, tail) {
  // Use substr to get the substring from the end of 'body' with length 'tail.length'
  let sub = body.substr(body.length - tail.length);

  // Check if the substring is equal to 'tail'
  return sub === tail;
}

// Example usage:
let result = correctTail("example", "ple");
console.log(result); // Output: true
