// https://www.codewars.com/kata/56f69d9f9400f508fb000ba7/train/javascript

function monkeyCount(n) {
  let result = [];
  for (i = 1; i <= n; i++) {
    result.push(i);
  }
  return result;
}

console.log(monkeyCount(2));
